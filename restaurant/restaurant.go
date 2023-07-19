package restaurant

import (
	"bytes"
	"encoding/binary"

	cmd "github.com/mikucat0309/mole/command"
	"github.com/mikucat0309/mole/game"
	map_ "github.com/mikucat0309/mole/map"
)

type (
	CookID uint32
)

type Info struct {
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
	User      game.User
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

var info *Info

type GetInfoData struct {
	User    uint32
	MapType map_.MapType
}

func GetInfo(c *game.GameConn) (info_ *Info, err error) {
	info = &Info{}
	resp, err := c.Conn.SendCmd(cmd.RESTAURANT_GET_INFO, GetInfoData{c.User, map_.RESTARUNT})
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
	info_ = info
	return
}
