package jsonbeego

type JsonResI interface {
	Error() *JsonRes
	ErrorWithErr(err error) *JsonRes
	ErrorWithStr(errMsg string) *JsonRes
	IsErr() bool
}

type JsonRes struct {
	Code Code        `jsonbeego:"code"`
	Msg  string      `jsonbeego:"msg"`
	Data interface{} `jsonbeego:"data"`
}

func (res *JsonRes) Error() *JsonRes {
	res.Code = Fail
	return res
}

func (res *JsonRes) ErrorWithStr(errMsg string) *JsonRes {
	res.Msg = errMsg
	return res
}

func (res *JsonRes) ErrorWithErr(err error) *JsonRes {
	return res.ErrorWithStr(err.Error())
}

func (res *JsonRes) IsErr() bool {
	return res.Code != Success
}

