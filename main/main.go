package main

import "fmt"
import "github.com/kataras/iris/v12"

func main() {
	fmt.Printf("main")
	app := iris.New()
	config := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,

		EnableOptimizations: false,
		Charset:             "UTF-8",
	})
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON("123213")
	})
	err := app.Run(iris.Addr(":8081"), config)
	if err != nil {
		panic("start server err:" + err.Error())
	}
}
