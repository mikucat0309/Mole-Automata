package login

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/mikucat0309/mole/conn"
)

type LoginReq struct {
	Pwd     [32]byte
	ZZZ1    int32
	ZZZ2    int32
	ZZZ3    int32
	Captcha [22]byte
}

func NewLoginReq(pwd [32]byte) *LoginReq {
	return &LoginReq{pwd, 0, 1, 0, [22]byte{}}
}

type LoginResp struct {
	Result  int32
	Session [16]byte
	Errlen  int32
}

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
	req := NewLoginReq(pwd)
	resp := &LoginResp{}
	err = c.SendRecv(CMD_LOGIN, req, resp)
	if err != nil {
		return [16]byte{}, errors.New(loginErrMsg[resp.Result])
	}
	c.Close()
	return resp.Session, nil
}
