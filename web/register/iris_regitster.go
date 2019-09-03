package register

import (
	"YTB-/config"
	"YTB-/web/middleware"
	"YTB-/web/router"
	"github.com/kataras/iris"
	recover2 "github.com/kataras/iris/middleware/recover"
)

func Iris() *iris.Application {

	var app *iris.Application

	if config.MyConfig.App.Debug {
		app = iris.Default()
	} else {
		app = iris.New()
		app.Use(recover2.New())
	}

	// 设置错误等级
	app.Logger().SetLevel(config.MyConfig.App.LogLevel)

	// 注册视图
	app.RegisterView(iris.HTML("./views", ".html"))

	// 注册全局流量控制
	app.Use(middleware.GlobalAfter)

	// 注册路由
	router.Register(app)

	return app
}
