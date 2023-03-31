package data

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
