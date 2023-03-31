package login

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/mikucat0309/mole/conn"
	"github.com/mikucat0309/mole/data"
)

const CMD_LOGIN = 103

var loginErrMsg = map[int32]string{
	0: "Success",
	1: "Incorrect password",
	2: "Incorrect captcha",
}

func md5text(pwd string) [32]byte {
	hash1 := []byte(fmt.Sprintf("%x", md5.Sum([]byte(pwd))))
	hash2 := []byte(fmt.Sprintf("%x", md5.Sum(hash1)))
	return *(*[32]byte)(hash2)
}

func Login(user uint32, pwd string) ([16]byte, error) {
	return LoginHash(user, md5text(pwd))
}

func LoginHash(user uint32, pwd [32]byte) ([16]byte, error) {
	c, err := conn.NewMoleConn("203.73.22.200:8888", user, false)
	if err != nil {
		return [16]byte{}, err
	}
	req := data.NewLoginReq(pwd)
	resp := &data.LoginResp{}
	err = c.SendRecv(CMD_LOGIN, req, resp)
	if err != nil {
		return [16]byte{}, errors.New(loginErrMsg[resp.Result])
	}
	c.Close()
	return resp.Session, nil
}
