package pkg

import "github.com/NageNalock/json-beego/jsonbeego"

type TestController struct {
	jbeego.JsonController
}

func (c *TestController) Get() {
	res := new(jbeego.JsonRes)
	defer c.Return(res)

	res.Data = "hello world"
}
