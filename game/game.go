package game

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"

	cmd "github.com/mikucat0309/mole/command"
	"github.com/mikucat0309/mole/conn"
)

type (
	User uint32
)

type GameConn struct {
	Conn *conn.MoleConn
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
		Conn: c,
		User: user,
	}
	_, err = c.SendCmd(cmd.LOGIN, NewGameLoginData(session))
	return c2, err
}
