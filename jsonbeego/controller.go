package jbeego

import (
	"github.com/beego/beego/v2/server/web"
)

type JsonController struct {
	web.Controller

	res JsonResI
}

func (c *JsonController) Return(res JsonResI) {
	c.res = res
	c.Data["json"] = c.res
	c.ServeJSON()
}
