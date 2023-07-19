package main

import (
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/mikucat0309/mole/game"
	"github.com/mikucat0309/mole/login"
	map_ "github.com/mikucat0309/mole/map"
	rest "github.com/mikucat0309/mole/restaurant"
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
	dishID := rest.DishID(util.UInt32(strconv.ParseUint(envs["DISH_ID"], 0, 32)))
	for {
		session, err := login.Login(user, pwd)
		if err != nil {
			log.Fatalf("登入失敗: %v\n", err)
		}
		c, err := game.GameLogin(user, session)
		if err != nil {
			log.Fatalf("登入伺服器失敗: %v\n", err)
		}
		interval := restaurant(c, dishID)
		log.Printf("等待 %d 秒", interval)
		time.Sleep(interval * time.Second)
	}
}

func roundRange(i uint32, d uint32) uint32 {
	r := i % d
	if r == 0 {
		return i
	}
	return i - r + d
}

var eventAnswers = map[rest.EventID]rest.EventAnswer{
	1:   1,
	2:   1,
	3:   2,
	4:   1,
	5:   2,
	6:   2,
	7:   2,
	8:   2,
	100: 1,
	200: 2,
	201: 3,
	202: 3,
	203: 3,
	204: 2,
	205: 3,
	206: 3,
	207: 3,
}

func restaurant(c *game.GameConn, dishID rest.DishID) time.Duration {
	interval := rest.DishInfos[dishID].CompleteDuration

	map_.LeaveMap(c)
	log.Println("離開地圖")

	map_.EnterMap(c, map_.RESTARUNT)
	log.Println("進入餐廳")

	info, err := rest.GetInfo(c)
	if err != nil {
		log.Fatalf("獲取餐廳資訊失敗: %v", err)
	}
	log.Printf("內部裝潢 %s, %d 個瓦斯爐, %d 個上菜盤, %d 個用餐位置", info.InnerStyle.Name, info.InnerStyle.Stove, info.InnerStyle.DishTable, info.InnerStyle.Table)

	event, err := rest.GetEvent(c)
	if err == nil && event != 0 {
		log.Printf("解決事件 %d", event)
		rest.SolveEvent(c, eventAnswers[event])
	}

	for _, stove := range info.Stoves {
		status := stove.Dish.Status()
		dishInfo := stove.Dish.Info()

		if status == rest.COOK_STATUS_PREPARE_1 || status == rest.COOK_STATUS_PREPARE_2 {
			rest.PrepareDish(c, &stove)
			log.Printf("瓦斯爐 %d 的 %s 開始料理, 剩餘 %d 秒", stove.Dish.Loc, dishInfo.Name, dishInfo.CompleteDuration)
			continue
		}

		if status == rest.COOK_STATUS_COOKING {
			remaining := dishInfo.CompleteDuration - stove.Dish.CookDuration
			log.Printf("瓦斯爐 %d 的 %s 正在料理, 剩餘 %d 秒", stove.Dish.Loc, dishInfo.Name, remaining)
			if roundRange(remaining, 60) < interval {
				interval = roundRange(remaining, 60)
			}
			continue
		}

		if status == rest.COOK_STATUS_COMPLETED {
			log.Printf("瓦斯爐 %d 的 %s 完成", stove.Dish.Loc, dishInfo.Name)
			err := rest.StoreDish(c, &stove)
			if err != nil {
				log.Printf("%v\n", err)
				continue
			}
		} else if status == rest.COOK_STATUS_EXPIRED {
			log.Printf("瓦斯爐 %d 的 %s 過期", stove.Dish.Loc, dishInfo.Name)
			rest.ClearDish(c, &stove)
		}

		dishInfo = rest.DishInfos[dishID]
		err := rest.MakeDish(c, &stove, dishID)
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}
		rest.PrepareDish(c, &stove)
		log.Printf("瓦斯爐 %d 的 %s 開始料理, 剩餘 %d 秒", stove.Dish.Loc, dishInfo.Name, dishInfo.CompleteDuration)
	}
	return time.Duration(interval)
}
