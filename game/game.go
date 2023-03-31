package game

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"log"

	"github.com/mikucat0309/mole/conn"
)

const (
	LOGIN              = 0xc9
	INIT_PLAYER        = 0x2b02
	WALK               = 0x12f
	LEAVE_MAP          = 0x192
	ENTER_MAP          = 0x191
	GET_RESTARUNT_INFO = 0x3f6
	GET_FOOD_EXP       = 0x416
	MAKE_FOOD          = 0x3f9
	PREPARE_FOOD       = 0x3fc
	STORE_FOOD         = 0x3fd
)

type GameConn struct {
	conn *conn.MoleConn
	User uint32
}

type LoginGameData struct {
	ZZZ1       int16
	Token      [16]byte
	SessionLen int32
	Session    [16]byte
	ZZZ2       int32
	ZZZ3       [64]byte
}

func NewGameLoginData(session [16]byte) *LoginGameData {
	tmp := fmt.Sprintf("%dhAo crAzy B%d", binary.BigEndian.Uint32(session[10:14]), binary.BigEndian.Uint32(session[3:7]))
	token := (*[16]byte)([]byte(fmt.Sprintf("%x", md5.Sum([]byte(tmp)))[6:22]))
	return &LoginGameData{1, *token, int32(len(session)), session, 0, [64]byte{0x30}}
}

func GameLogin(user uint32, session [16]byte) (*GameConn, error) {
	c, err := conn.NewMoleConn("203.73.22.191:1201", user, true)
	if err != nil {
		return nil, err
	}
	c2 := &GameConn{
		conn: c,
		User: user,
	}
	_, err = c.SendCmd(LOGIN, NewGameLoginData(session))
	if err != nil {
		return nil, err
	}
	return c2, nil
}

func (c *GameConn) LeaveMap() {
	c.conn.SendCmd(LEAVE_MAP, []byte{})
}

type EnterMapData struct {
	User uint32
	Map  int32
	ZZZ1 int32
	ZZZ2 int32
	ZZZ3 int32
	ZZZ4 int32
}

func (c *GameConn) EnterMap(mapID int32) {
	c.conn.SendCmd(ENTER_MAP, EnterMapData{c.User, mapID, 0, 0, 0, 0})
}

// Restarunt

type (
	Dish     int32
	Cook     int32
	CookLoc  int32
	StoreLoc int32
)

type MakeDishData struct {
	Dish Dish
	Loc  CookLoc
}

type MakeDishRespData struct {
	Dish     Dish
	Cook     Cook
	CookLoc  CookLoc
	CookProg int32
	ZZZ1     int32
}

type PrepareDishData struct {
	Dish Dish
	Cook Cook
}

func (c *GameConn) MakeDish(dish Dish, loc CookLoc) (Cook, error) {
	resp := &MakeDishRespData{}
	err := c.conn.SendRecv(MAKE_FOOD, MakeDishData{dish, loc}, resp)
	if err != nil {
		return 0, err
	}
	return resp.Cook, nil
}

func (c *GameConn) PrepareDish(dish Dish, cook Cook) {
	_, err := c.conn.SendCmd(PREPARE_FOOD, PrepareDishData{dish, cook})
	if err != nil {
		log.Println(err)
	}
}

type StoreDishData struct {
	Dish     Dish
	Cook     Cook
	CookLoc  CookLoc
	StoreLoc StoreLoc
}

func (c *GameConn) StoreDish(dish Dish, cook Cook, cookLoc CookLoc, storeLoc StoreLoc) {
	_, err := c.conn.SendCmd(STORE_FOOD, StoreDishData{dish, cook, cookLoc, storeLoc})
	if err != nil {
		log.Println(err)
	}
}
