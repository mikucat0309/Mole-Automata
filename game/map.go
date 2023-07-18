package game

type (
	MapType uint32
)

const (
	MAP_FARM        MapType = 0x2
	MAP_RESTARUNT   MapType = 0x1f
	MAP_LAHM        MapType = 0x20
	MAP_PIG         MapType = 0x22
	MAP_PIG_BEAUTY  MapType = 0x23
	MAP_PIG_FACTORY MapType = 0x24
)

func (c *GameConn) LeaveMap() {
	c.conn.SendCmd(CMD_LEAVE_MAP, []byte{})
}

type EnterMapData struct {
	MapID      uint32
	MapType    MapType
	OldMapID   int32
	OldMapType int32
	NewGrid    int32
	OldGrid    int32
}

func (c *GameConn) EnterMap(mapType MapType) {
	c.conn.SendCmd(CMD_ENTER_MAP, EnterMapData{c.User, mapType, 0, 0, 0, 0})
}
