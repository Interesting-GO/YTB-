package router

import (
	"YTB-/web/controller"
	"github.com/kataras/iris"
)

func Register(app *iris.Application) {
	app.Post("/download", controller.Download)

	app.Get("/dow-data", controller.GetData)
}
