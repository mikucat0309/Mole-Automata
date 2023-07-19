package map_

import (
	cmd "github.com/mikucat0309/mole/command"
	"github.com/mikucat0309/mole/game"
)

type (
	MapType uint32
)

const (
	FARM        MapType = 0x2
	RESTARUNT   MapType = 0x1f
	LAHM        MapType = 0x20
	PIG         MapType = 0x22
	PIG_BEAUTY  MapType = 0x23
	PIG_FACTORY MapType = 0x24
)

func LeaveMap(c *game.GameConn) {
	c.Conn.SendCmd(cmd.MAP_LEAVE, []byte{})
}

type EnterMapData struct {
	MapID      uint32
	MapType    MapType
	OldMapID   int32
	OldMapType int32
	NewGrid    int32
	OldGrid    int32
}

func EnterMap(c *game.GameConn, mapType MapType) {
	c.Conn.SendCmd(cmd.MAP_ENTER, EnterMapData{c.User, mapType, 0, 0, 0, 0})
}
