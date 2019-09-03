package main

import (
	"YTB-/config"
	"YTB-/web/register"
	"github.com/kataras/iris"
)

func main() {
	app := register.Iris()

	app.Run(iris.Addr(config.MyConfig.App.Host),iris.WithCharset("UTF-8"))
}

