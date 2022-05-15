package jsonbeego

import (
	"github.com/beego/beego/v2/server/web"
)

type JsonController struct {
	web.Controller

	res JsonResI
}

func (c *JsonController) Return(res JsonResI) {
	c.res = res
	c.Data["jsonbeego"] = c.res
	c.ServeJSON()
}
