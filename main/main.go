package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/rpc"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Time int64  `json:"time"`
}

func main() {
	app := iris.New()
	config := iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,

		EnableOptimizations: false,
		Charset:             "UTF-8",
	})
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON("123213")
	})

	conn, err := rpc.Dial("tcp", "localhost:9999") //发起连接
	if err != nil {
		fmt.Println("dail err")
	}

	//defer conn.Close()
	var reply string
	app.Get("/push", func(ctx iris.Context) {
		err = conn.Call("hello.Add", "参数", &reply) //进行rpc调用
		if err != nil {
			fmt.Printf("call err" + err.Error())

		}
		fmt.Println(reply)

		ctx.JSON(reply)
	})
	app.Get("/test", func(ctx iris.Context) {
		err = conn.Call("hello.HelloWorld", "xf", &reply) //进行rpc调用
		if err != nil {
			fmt.Printf("call err" + err.Error())
		}
		fmt.Println(reply)
		ctx.JSON(reply)
	})
	app.Get("/test1", func(ctx iris.Context) {
		user := User{}
		err = conn.Call("hello.GetUser", User{"zhansan", 12, 0}, &user) //进行rpc调用
		if err != nil {
			fmt.Printf("call err" + err.Error())
		}
		fmt.Println(user)
		ctx.JSON(user)
	})
	err = app.Run(iris.Addr(":8081"), config)
	if err != nil {
		panic("start server err:" + err.Error())
	}
}
