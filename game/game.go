package game

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"

	"github.com/mikucat0309/mole/conn"
)

type (
	User uint32
)

const (
	CMD_LOGIN              = 0xc9
	CMD_INIT_PLAYER        = 0x2b02
	CMD_WALK               = 0x12f
	CMD_LEAVE_MAP          = 0x192
	CMD_ENTER_MAP          = 0x191
	CMD_GET_RESTARUNT_INFO = 0x3f6
	CMD_GET_FOOD_EXP       = 0x416
	CMD_MAKE_FOOD          = 0x3f9
	CMD_CLEAR_FOOD         = 0x3fb
	CMD_PREPARE_FOOD       = 0x3fc
	CMD_STORE_FOOD         = 0x3fd
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
	_, err = c.SendCmd(CMD_LOGIN, NewGameLoginData(session))
	return c2, err
}
