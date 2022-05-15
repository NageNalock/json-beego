package example

import "github.com/NageNalock/json-beego/jsonbeego"

type TestController struct {
	jsonbeego.JsonController
}

func (c *TestController) Get() {
	res := new(jsonbeego.JsonRes)
	defer c.Return(res)

	res.Data = "hello world"
}
