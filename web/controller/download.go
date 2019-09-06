package controller

import (
	"YTB-/config"
	"YTB-/datamodels"
	"YTB-/datasource/pgsql_conn"
	"YTB-/defs"
	"YTB-/resp"
	"github.com/Interesting-GO/youtubetools/video_dow"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/kataras/iris"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	num      int
	dataChan chan *defs.YouTuBeRsq
	dataMax  chan int
)

func init() {
	dataChan = make(chan *defs.YouTuBeRsq, config.MyConfig.App.TaskNum)
	dataMax = make(chan int, config.MyConfig.App.TaskNum)
	log.Print("队列初始化成功！")
	num = 0
}

func Download(ctx iris.Context) {
	input := defs.YouTuBeRsq{}
	err := ctx.ReadForm(&input)
	if err != nil {
		log.Println(err.Error())
		resp.Resp(ctx, defs.TaskErrorReq)
		return
	}

	clog.Println(input)

	// 判断参数是否错误
	if strings.Index(input.Url, "https://") == -1 {
		resp.Resp(ctx, defs.TaskErrorReq)
		return
	}

	// 查询数据是否存在如果存在则跳过
	exit := GetDataExit(input.Id)
	if !exit {
		// 如果存在就返回任务以存在
		resp.Resp(ctx, defs.Task{Code: 400, Msg: "任务以存在"})
		return
	}

	// 如果存在就加入下载队列

	// 判断队伍是否满载
	if len(dataMax) >= config.MyConfig.App.TaskNum {
		// 队伍满载 返回繁忙信息
		resp.Resp(ctx, defs.TaskError)
		return
	} else {
		dataChan <- &input
		dataMax <- 1

		resp.Resp(ctx, defs.TaskOk)
	}

}

func GetData(ctx iris.Context) {
	// 向数据库查询数据
	var datas []*datamodels.Video
	_, e := pgsql_conn.PgDb.QueryTable("video").All(&datas)
	if e != nil {
		ctx.StatusCode(500)
		ctx.JSON("数据查询错误")
	}
	ctx.JSON(datas)
}

func DowTask() {
	for {
		select {
		case data := <-dataChan:
			// 开启下载协程
			go dow(data)
		}
	}
}

// 下载阶段
func dow(data *defs.YouTuBeRsq) {
	// 获取当前时间  生成目录地址
	s, e := easyutils.TimeGetTimeToString(easyutils.TimeGetNowTimeStr())
	if e != nil {
		panic(e.Error())
	}

	name := easyutils.SuperRand()
	path := "./video/" + s + "/"
	pathurl := "/" + s + "/"

	e = easyutils.DirPing(path)
	if e != nil {
		panic(e.Error())
	}

	// 下载文件

	// 失败尝试 10 次 每次 休息 10s

	k := 0
	for {
		e := video_dow.YoutubeDow(data.Url, path+name+".mp4", "127.0.0.1:8001")
		if e != nil {
			k += 1
			if k < 10 {
				clog.Println("=================下载失败进行尝试   " + strconv.Itoa(k))
				log.Println(e.Error())
				time.Sleep(time.Second * 10)
				continue
			} else {
				// 入库标记下载失败
				data := datamodels.Video{VideoId: data.Id, Name: data.Name, Path: e.Error()}
				_, e := pgsql_conn.PgDb.Insert(&data)
				if e != nil {
					clog.Println("sql 生成错误")
				}
				return
			}

		} else {
			time.Sleep(3 * time.Second)
			break
		}
	}
	// 下载完毕 入库

	datas := datamodels.Video{VideoId: data.Id, Name: data.Name, Path: pathurl + name + ".mp4"}
	_, e = pgsql_conn.PgDb.Insert(&datas)
	if e != nil {
		clog.Println("sql 生成错误")
	}
	num += 1
	clog.Println("下载成功   第: " + strconv.Itoa(num))
	<-dataMax
}

func GetDataExit(id string) bool {
	data := new(datamodels.Video)
	table := pgsql_conn.PgDb.QueryTable(data)
	exist := table.Filter("video_id", id).Exist()
	log.Println(exist)
	return !exist
}
