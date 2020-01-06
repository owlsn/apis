package main

import (
	"flag"
	"os"
	"github.com/owlsn/apis/src/common/config"
	"github.com/sirupsen/logrus"
	"github.com/owlsn/apis/src/app"
)

var configFile = flag.String("config", "./server.yaml", "config file path")

func init() {
	flag.Parse()

<<<<<<< HEAD
	app.Handle("POST", "/api/index", func(ctx iris.Context) {
		// {"title":"hello", "content":"helloworld"}
		ctx.JSON(iris.Map{
			"title":   "/api/index",
			"content": "helloworld",
		})
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	// app.Get("/ping", func(ctx iris.Context) {
	// 	ctx.WriteString("pong")
	// })

	// // Method:   GET
	// // Resource: http://localhost:8080/hello
	// app.Get("/hello", func(ctx iris.Context) {
	// 	ctx.JSON(iris.Map{"message": "Hello Iris!"})
	// })
=======
	config.InitConfig(*configFile) 
	// 初始化配置
	file, err := os.OpenFile(config.Conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Error(err)
	}
}
>>>>>>> dev_owlsn

func main() {
	app.StartOn()
	app.InitIris()
}
