package pgsql_conn

import (
	"YTB-/config"
	"YTB-/datamodels"
	"github.com/dollarkillerx/beegoorm"
	_ "github.com/lib/pq"
)

var (
	err  error
)

func init() {

	err = beegoorm.RegisterDataBase("default", "postgres",
		config.MyConfig.Pgsql.Dsn)
	if err != nil {
		panic(err.Error())
	}

	beegoorm.SetMaxOpenConns("default", config.MyConfig.Pgsql.MaxOpen)
	beegoorm.SetMaxIdleConns("default", config.MyConfig.Pgsql.MaxIdle)

	if config.MyConfig.App.Debug {
		beegoorm.Debug = true
	}

	mapping()




}

func PgDb() beegoorm.Ormer {
	PgDb := beegoorm.NewOrm()
	return PgDb
}

// 数据库映射
func mapping() {
	// register model  注册模型
	beegoorm.RegisterModel(new(datamodels.Video))

	// 完成映射
	err = beegoorm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err.Error())
	}
}
