package main

import (
	"flag"
	"os"
	"github.com/owlsn/apis/src/common/config"
	"github.com/sirupsen/logrus"
	"github.com/owlsn/apis/src/app"
	"github.com/owlsn/apis/src/utils/database"
	"reflect"
)

var configFile = flag.String("config", "./server.yaml", "config file path")

func init() {
	flag.Parse()

	config.InitConfig(*configFile) 
	// 初始化配置
	file, err := os.OpenFile(config.Conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Error(err)
	}

	db, err := database.Instance() // 连接数据库
	if reflect.TypeOf(db).String() != "*gorm.DB" || err != nil{
		logrus.Errorf("get gorm.DB fail " , err.Error())
	}else{
		// defer db.Close()
	}
}

func main() {
	app.StartOn()
	app.InitIris()
}
