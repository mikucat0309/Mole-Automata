package restaurant

import (
	cmd "github.com/mikucat0309/mole/command"
	"github.com/mikucat0309/mole/game"
)

type (
	EventID     uint32
	EventAnswer uint32
)

func GetEvent(c *game.GameConn) (eventID EventID, err error) {
	err = c.Conn.SendRecv(cmd.RESTAURUNT_GET_EVENT, []byte{}, &eventID)
	return
}

func SolveEvent(c *game.GameConn, ans EventAnswer) (err error) {
	_, err = c.Conn.SendCmd(cmd.RESTAURUNT_SOLVE_EVENT, ans)
	return
}
