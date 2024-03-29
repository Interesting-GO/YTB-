package config

import (
	"YTB-/defs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type myconf struct {
	App struct {
		Host       string `yaml:"host"`
		Debug      bool   `yaml:"debug"`
		MaxRequest int    `yaml:"max_request"`
		LogLevel   string `yaml:"log_level"`
		TaskNum    int    `yaml:"task_num"`
	}
	Mysql struct {
		Dsn   string `yaml:"dsn"`
		Cache bool   `yaml:"cache"`
	}
	Pgsql struct {
		Dsn     string        `yaml:"dsn"`
		MaxIdle int           `yaml:"max_idle"`
		MaxOpen int           `yaml:"max_open"`
		TimeOut time.Duration `yaml:"time_out"`
	}
	Redis struct {
		Maxidle     int           `yaml:"maxidle"`
		MaxActive   int           `yaml:"max_active"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
		Port        string        `yaml:"port"`
	}
}

var (
	MyConfig *myconf

	DataChan chan *defs.YouTuBeRsq
	DataMax  chan int
)

func init() {
	MyConfig = &myconf{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, MyConfig)
	if e != nil {
		panic(e.Error())
	}

	if MyConfig.App.Debug {
		log.Println(MyConfig)
	}

	DataChan = make(chan *defs.YouTuBeRsq, MyConfig.App.TaskNum)
	DataMax = make(chan int, MyConfig.App.TaskNum)
}
