package restaurant

import (
	"fmt"

	cmd "github.com/mikucat0309/mole/command"
	"github.com/mikucat0309/mole/game"
)

type (
	DishContainerLoc uint32
	CookStatus       uint32
)

type DishContainer struct {
	CookID       CookID
	Loc          DishContainerLoc
	Dish         DishID
	Count        uint32
	state        uint32
	CookDuration uint32
}

const (
	COOK_STATUS_EMPTY     CookStatus = 0
	COOK_STATUS_PREPARE_1 CookStatus = 1
	COOK_STATUS_PREPARE_2 CookStatus = 2
	COOK_STATUS_COOKING   CookStatus = 3
	COOK_STATUS_COMPLETED CookStatus = 4
	COOK_STATUS_EXPIRED   CookStatus = 5
)

type Stove struct {
	Dish *DishContainer
}

type DishTable struct {
	Dish *DishContainer
}

func (d *DishContainer) isEmpty() bool {
	return d.Count == 0
}

func (d *DishContainer) Info() *DishInfo {
	return DishInfos[d.Dish]
}

func (d *DishContainer) Status() CookStatus {
	if d.isEmpty() {
		return COOK_STATUS_EMPTY
	}
	info := d.Info()
	if d.state == 1 {
		return COOK_STATUS_PREPARE_1
	}
	if d.state == 2 {
		return COOK_STATUS_PREPARE_2
	}
	if d.CookDuration < info.CompleteDuration {
		return COOK_STATUS_COOKING
	}
	if d.CookDuration < info.CompleteDuration+info.ExpireDuration {
		return COOK_STATUS_COMPLETED
	}
	return COOK_STATUS_EXPIRED
}

type MakeDishData struct {
	Dish  DishID
	Stove DishContainerLoc
}

type MakeDishRespData struct {
	Dish     DishID
	Cook     CookID
	Stove    DishContainerLoc
	CookProg int32
	ZZZ1     int32
}

func MakeDish(c *game.GameConn, stove *Stove, dish DishID) (err error) {
	resp := &MakeDishRespData{}
	err = c.Conn.SendRecv(cmd.RESTAURANT_MAKE_FOOD, MakeDishData{dish, stove.Dish.Loc}, resp)
	stove.Dish.Dish = dish
	stove.Dish.CookID = resp.Cook
	stove.Dish.state = 1
	return
}

type PrepareDishData struct {
	Dish DishID
	Cook CookID
}

func PrepareDish(c *game.GameConn, stove *Stove) (err error) {
	if stove.Dish.state == 1 {
		_, err = c.Conn.SendCmd(cmd.RESTAURANT_PREPARE_FOOD, PrepareDishData{stove.Dish.Dish, stove.Dish.CookID})
		stove.Dish.state = 2
	}
	if stove.Dish.state == 2 {
		_, err = c.Conn.SendCmd(cmd.RESTAURANT_PREPARE_FOOD, PrepareDishData{stove.Dish.Dish, stove.Dish.CookID})
		stove.Dish.state = 3
		stove.Dish.CookDuration = 0
	}
	return
}

type StoreDishData struct {
	Dish      DishID
	Cook      CookID
	Stove     DishContainerLoc
	DishTable DishContainerLoc
}

func StoreDish(c *game.GameConn, stove *Stove) (err error) {
	table := findTargetDishtable(info.DishTables, stove.Dish.Dish)
	if table == nil {
		return fmt.Errorf("沒有多餘的上菜盤可以放置新料理")
	}
	_, err = c.Conn.SendCmd(cmd.RESTAURANT_STORE_FOOD, StoreDishData{stove.Dish.Dish, stove.Dish.CookID, stove.Dish.Loc, table.Dish.Loc})
	if err != nil {
		return
	}
	stove.Dish = &DishContainer{0, stove.Dish.Loc, 0, 0, 0, 0}
	return
}

func findTargetDishtable(tables []DishTable, dish DishID) *DishTable {
	for k := range tables {
		if tables[k].Dish.Dish == dish {
			return &tables[k]
		}
	}
	for k := range tables {
		if tables[k].Dish.Status() == COOK_STATUS_EMPTY {
			return &tables[k]
		}
	}
	return nil
}

type ClearDishData struct {
	Dish  DishID
	Cook  CookID
	Stove DishContainerLoc
}

func ClearDish(c *game.GameConn, stove *Stove) (err error) {
	_, err = c.Conn.SendCmd(cmd.RESTAURANT_CLEAR_FOOD, ClearDishData{stove.Dish.Dish, stove.Dish.CookID, stove.Dish.Loc})
	if err != nil {
		return
	}
	stove.Dish = &DishContainer{0, stove.Dish.Loc, 0, 0, 0, 0}
	return
}
