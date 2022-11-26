package main

import (
	"fmt"
	"github.com/NageNalock/json-beego/example/pkg"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	// config file is invalid for example
	//err := jbeego.InitLog()
	//if err != nil {
	//	panic(err.Error())
	//}
	//port := config.DefaultInt("port", 3306)

	pkg.InitRouter()

	logs.Info("example start")

	web.Run(fmt.Sprintf(":%d", 3306))
}
