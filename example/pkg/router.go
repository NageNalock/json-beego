package pkg

import (
	jbeego "github.com/NageNalock/json-beego/jsonbeego"
	"github.com/beego/beego/v2/server/web"
)

func InitRouter() {
	web.AddNamespace(web.NewNamespace("/v1",
		web.NSRouter("/hello", &TestController{}, jbeego.Get+"Get"),
	))
}
