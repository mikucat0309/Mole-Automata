package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/mikucat0309/mole/game"
	"github.com/mikucat0309/mole/login"
)

func main() {
	var (
		err         error
		u64         uint64
		user        uint32
		pwd         string
		stoveAmount int
		dish        game.Dish = 0x147261
	)

	var envs map[string]string
	envs, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}
	fmt.Printf("User: %s\n", envs["USER"])
	u64, _ = strconv.ParseUint(envs["USER"], 0, 32)
	user = uint32(u64)
	pwd = envs["PASSWORD"]
	u64, _ = strconv.ParseUint(envs["STOVE_AMOUNT"], 0, 32)
	stoveAmount = int(u64)
	u64, _ = strconv.ParseUint(envs["DISH_ID"], 0, 32)
	dish = game.Dish(u64)

	session, err := login.Login(user, pwd)
	if err != nil {
		log.Printf("Login error: %v\n", err)
		return
	}
	fmt.Printf("Session: %x\n", session)
	c, err := game.GameLogin(user, session)
	if err != nil {
		log.Printf("Login game error: %v\n", err)
		return
	}
	c.LeaveMap()
	fmt.Println("離開地圖")

	c.EnterMap(0x1f)
	fmt.Println("進入地圖")

	cooks := make([]game.Cook, 10)
	for i := 1; i <= stoveAmount; i++ {
		cooks[i], err = c.MakeDish(dish, game.CookLoc(i))
		if err != nil {
			break
		}
		c.PrepareDish(dish, cooks[i])
		c.PrepareDish(dish, cooks[i])
		fmt.Printf("瓦斯爐 %d 成功烹煮料理 %d，烹煮序號: %d\n", i, dish, cooks[i])
		time.Sleep(1 * time.Second)
	}
}
