package main

import (
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	// app.Handle("GET", "/index", func(ctx iris.Context){
	// 	ctx.HTML("<h1>index</h1>")
	// })

	app.Handle("POST", "/index", func(ctx iris.Context) {
		// {"title":"hello", "content":"helloworld"}
		ctx.JSON(iris.Map{
			"title":   "index",
			"content": "helloworld",
		})
	})

	app.Handle("POST", "/", func(ctx iris.Context) {
		// {"title":"hello", "content":"helloworld"}
		ctx.JSON(iris.Map{
			"title":   "/",
			"content": "helloworld",
		})
	})

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

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
