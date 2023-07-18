package main

import (
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/mikucat0309/mole/game"
	"github.com/mikucat0309/mole/login"
	"github.com/mikucat0309/mole/util"
)

func main() {
	envs, err := godotenv.Read()
	if err != nil {
		log.Fatal("讀取 .env 檔案失敗")
	}
	log.Printf("帳號: %s\n", envs["USER"])
	user := util.UInt32(strconv.ParseUint(envs["USER"], 0, 32))
	pwd := envs["PASSWORD"]
	dishID := game.DishID(util.UInt32(strconv.ParseUint(envs["DISH_ID"], 0, 32)))
	for {
		interval := mainloop(user, pwd, dishID)
		log.Printf("等待 %d 秒", interval)
		time.Sleep(interval * time.Second)
	}
}

func RoundRange(i uint32, d uint32) uint32 {
	r := i % d
	if r == 0 {
		return i
	}
	return i - r + d
}

func mainloop(user uint32, pwd string, dishID game.DishID) time.Duration {
	interval := game.DishInfos[dishID].CompleteDuration
	session, err := login.Login(user, pwd)
	if err != nil {
		log.Fatalf("登入失敗: %v\n", err)
	}
	c, err := game.GameLogin(user, session)
	if err != nil {
		log.Fatalf("登入伺服器失敗: %v\n", err)
	}
	c.LeaveMap()
	log.Println("離開地圖")

	c.EnterMap(game.MAP_RESTARUNT)
	log.Println("進入餐廳")

	info, err := c.GetRestaruntInfo()
	if err != nil {
		log.Fatalf("獲取餐廳資訊失敗: %v", err)
	}
	log.Printf("內部裝潢 %s, %d 個瓦斯爐, %d 個上菜盤, %d 個用餐位置", info.InnerStyle.Name, info.InnerStyle.Stove, info.InnerStyle.DishTable, info.InnerStyle.Table)
	for _, stove := range info.Stoves {
		status := stove.Dish.Status()
		dishInfo := stove.Dish.Info()

		if status == game.COOK_STATUS_PREPARE_1 || status == game.COOK_STATUS_PREPARE_2 {
			log.Printf("瓦斯爐 %d 的 %s 開始料理, 剩餘 %d 秒", stove.Dish.Loc, dishInfo.Name, dishInfo.CompleteDuration)
			c.PrepareDish(&stove)
			status = game.COOK_STATUS_COOKING
			continue
		}

		if status == game.COOK_STATUS_COOKING {
			remaining := dishInfo.CompleteDuration - stove.Dish.CookDuration
			log.Printf("瓦斯爐 %d 的 %s 正在料理, 剩餘 %d 秒", stove.Dish.Loc, dishInfo.Name, remaining)
			if RoundRange(remaining, 60) < interval {
				interval = RoundRange(remaining, 60)
			}
			continue
		}

		if status == game.COOK_STATUS_COMPLETED {
			log.Printf("瓦斯爐 %d 的 %s 完成", stove.Dish.Loc, dishInfo.Name)
			err := c.StoreDish(&stove)
			if err != nil {
				log.Printf("%v\n", err)
				continue
			}
		} else if status == game.COOK_STATUS_EXPIRED {
			log.Printf("瓦斯爐 %d 的 %s 過期", stove.Dish.Loc, dishInfo.Name)
			c.ClearDish(&stove)
		}

		dishInfo = game.DishInfos[dishID]
		log.Printf("瓦斯爐 %d 的 %s 開始料理, 剩餘 %d 秒", stove.Dish.Loc, dishInfo.Name, dishInfo.CompleteDuration)
		err := c.MakeDish(&stove, dishID)
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}
		c.PrepareDish(&stove)
	}
	return time.Duration(interval)
}
