package pgsql_conn

import (
	"YTB-/config"
	"github.com/dollarkillerx/beegoorm"
)

var (
	PgDb beegoorm.Ormer
	err  error
)

func init() {

	err = beegoorm.RegisterDataBase("default", "postgres",
		"user=navi password=psql233 dbname=navi host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	beegoorm.SetMaxOpenConns("default", config.MyConfig.Pgsql.MaxOpen)
	beegoorm.SetMaxIdleConns("default", config.MyConfig.Pgsql.MaxIdle)

	PgDb = beegoorm.NewOrm()

	mapping()

}

// 数据库映射
func mapping() {


	// register model  注册模型
	beegoorm.RegisterModel()

	// 完成映射
	err = beegoorm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err.Error())
	}
}
