package resp

import (
	"YTB-/defs"
	"github.com/kataras/iris"
)

func Resp(ctx iris.Context, data defs.Task) {
	ctx.StatusCode(data.Code)
	ctx.JSON(data)
}
