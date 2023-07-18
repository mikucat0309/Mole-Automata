package game

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type (
	CookID uint32
)

type GetRestaruntInfoData struct {
	User    uint32
	MapType MapType
}

type RestaruntInfo struct {
	User       uint32
	Serial     uint32
	Grid       uint32
	Exp        uint32
	Money      uint32
	Favor      uint32
	Type       uint32
	Level      uint32
	AllFood    uint32
	InnerStyle *InnerStyle
	Style      uint32
	Name       string
	PeopMoney  uint32
	StarCount  uint32
	Stoves     []Stove
	DishTables []DishTable
	Employees  []Employee
}

type Employee struct {
	User      User
	ID        uint32
	Name      string
	Color     uint32
	Level     uint32
	Skill     uint32
	EmpLevel  uint32
	EmpMoney  uint32
	EndTime   uint32
	TimeLimit uint32
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

type PrepareDishData struct {
	Dish DishID
	Cook CookID
}

type StoreDishData struct {
	Dish      DishID
	Cook      CookID
	Stove     DishContainerLoc
	DishTable DishContainerLoc
}

type ClearDishData struct {
	Dish  DishID
	Cook  CookID
	Stove DishContainerLoc
}

var restaruntInfo *RestaruntInfo

func (c *GameConn) MakeDish(stove *Stove, dish DishID) (err error) {
	resp := &MakeDishRespData{}
	err = c.conn.SendRecv(CMD_MAKE_FOOD, MakeDishData{dish, stove.Dish.Loc}, resp)
	stove.Dish.Dish = dish
	stove.Dish.CookID = resp.Cook
	stove.Dish.state = 1
	return
}

func (c *GameConn) PrepareDish(stove *Stove) (err error) {
	if stove.Dish.state == 1 {
		_, err = c.conn.SendCmd(CMD_PREPARE_FOOD, PrepareDishData{stove.Dish.Dish, stove.Dish.CookID})
		stove.Dish.state = 2
	}
	if stove.Dish.state == 2 {
		_, err = c.conn.SendCmd(CMD_PREPARE_FOOD, PrepareDishData{stove.Dish.Dish, stove.Dish.CookID})
		stove.Dish.state = 3
		stove.Dish.CookDuration = 0
	}
	return
}

func (c *GameConn) StoreDish(stove *Stove) error {
	table := findTargetDishtable(restaruntInfo.DishTables, stove.Dish.Dish)
	if table == nil {
		return fmt.Errorf("沒有多餘的上菜盤可以放置新料理")
	}
	_, err := c.conn.SendCmd(CMD_STORE_FOOD, StoreDishData{stove.Dish.Dish, stove.Dish.CookID, stove.Dish.Loc, table.Dish.Loc})
	if err != nil {
		return err
	}
	stove.Dish = &DishContainer{0, stove.Dish.Loc, 0, 0, 0, 0}
	return nil
}

func findTargetDishtable(tables []DishTable, dish DishID) *DishTable {
	for k := range tables {
		if tables[k].Dish.Dish == dish {
			return &tables[k]
		}
	}
	for k := range tables {
		if tables[k].Dish.IsEmpty() {
			return &tables[k]
		}
	}
	return nil
}

func (c *GameConn) ClearDish(stove *Stove) error {
	_, err := c.conn.SendCmd(CMD_CLEAR_FOOD, ClearDishData{stove.Dish.Dish, stove.Dish.CookID, stove.Dish.Loc})
	if err != nil {
		return err
	}
	return nil
}

func (c *GameConn) GetRestaruntInfo() (*RestaruntInfo, error) {
	info := &RestaruntInfo{}
	resp, err := c.conn.SendCmd(CMD_GET_RESTARUNT_INFO, GetRestaruntInfoData{c.User, MAP_RESTARUNT})
	reader := bytes.NewReader(resp)

	var innerStyleId InnerStyleId
	strBuf := make([]byte, 16)
	var count uint32

	binary.Read(reader, binary.BigEndian, &info.User)
	binary.Read(reader, binary.BigEndian, &info.Serial)
	binary.Read(reader, binary.BigEndian, &info.Grid)
	binary.Read(reader, binary.BigEndian, &info.Exp)
	binary.Read(reader, binary.BigEndian, &info.Money)
	binary.Read(reader, binary.BigEndian, &info.Favor)
	binary.Read(reader, binary.BigEndian, &info.Type)
	binary.Read(reader, binary.BigEndian, &info.Level)
	binary.Read(reader, binary.BigEndian, &info.AllFood)
	binary.Read(reader, binary.BigEndian, &innerStyleId)
	info.InnerStyle = InnerStyles[innerStyleId]
	binary.Read(reader, binary.BigEndian, &info.Style)
	binary.Read(reader, binary.BigEndian, strBuf)
	info.Name = string(strBuf)
	binary.Read(reader, binary.BigEndian, &info.PeopMoney)
	binary.Read(reader, binary.BigEndian, &info.StarCount)

	info.Stoves = make([]Stove, info.InnerStyle.Stove)
	info.DishTables = make([]DishTable, info.InnerStyle.DishTable)
	for k := range info.Stoves {
		info.Stoves[k].Dish = &DishContainer{0, DishContainerLoc(k + 1), 0, 0, 0, 0}
	}
	for k := range info.DishTables {
		info.DishTables[k].Dish = &DishContainer{0, DishContainerLoc(k + 51), 0, 0, 0, 0}
	}
	binary.Read(reader, binary.BigEndian, &count)
	for i := 0; i < int(count); i++ {
		container := &DishContainer{}
		binary.Read(reader, binary.BigEndian, &container.Loc)
		binary.Read(reader, binary.BigEndian, &container.Dish)
		binary.Read(reader, binary.BigEndian, &container.CookID)
		binary.Read(reader, binary.BigEndian, &container.Count)
		binary.Read(reader, binary.BigEndian, &container.state)
		binary.Read(reader, binary.BigEndian, &container.CookDuration)
		if container.Loc < 10 {
			info.Stoves[container.Loc-1].Dish = container
		} else {
			info.DishTables[container.Loc-51].Dish = container
		}
	}

	binary.Read(reader, binary.BigEndian, &count)
	employees := make([]Employee, count)
	for i := 0; i < int(count); i++ {
		binary.Read(reader, binary.BigEndian, &employees[i].User)
		binary.Read(reader, binary.BigEndian, &employees[i].ID)
		binary.Read(reader, binary.BigEndian, strBuf)
		employees[i].Name = string(strBuf)
		binary.Read(reader, binary.BigEndian, &employees[i].Color)
		binary.Read(reader, binary.BigEndian, &employees[i].Level)
		binary.Read(reader, binary.BigEndian, &employees[i].Skill)
		binary.Read(reader, binary.BigEndian, &employees[i].EmpLevel)
		binary.Read(reader, binary.BigEndian, &employees[i].EmpMoney)
		binary.Read(reader, binary.BigEndian, &employees[i].EndTime)
		binary.Read(reader, binary.BigEndian, &employees[i].TimeLimit)
	}
	info.Employees = employees
	restaruntInfo = info
	return restaruntInfo, err
}
