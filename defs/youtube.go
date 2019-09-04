package defs

type YouTuBeRsq struct {
	Id   string `json:"id" form:"id"`     // 视频id
	Name string `json:"name" form:"name"` // 动漫名称
	Url  string `json:"url" form:"url"`   // url地址
}

type M3U8Success struct {
	Id   string `json:"id" form:"id"`     // 视频id
	Path string `json:"path" form:"path"` // 视频路径
	Msg  string `json:"msg" form:"msg"`   // 信息
	Code int    `json:"code" form:"code"` // 判断码  200正常 500错误
}

type Task struct { // 任务下发时返回信息
	Code int    `json:"code"` // 代码 200正确 503服务器任务队列已满(客户端需等待)
	Msg  string `json:"msg"`  // 消息
}

var (
	TaskOk       = Task{Code: 200, Msg: "任务下发完成"}
	TaskError    = Task{Code: 503, Msg: "任务队列已满"}
	TaskErrorReq = Task{Code: 400, Msg: "参数错误"}
)
