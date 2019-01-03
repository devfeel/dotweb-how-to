* How To:
  - how to query slow response in your app?
  ~~~go
  func main() {
  	app := dotweb.Classic("/home/logs/wwwroot/")

  	// deal slow response, use default handler
  	// default Handler will write timeout-response in dotweb's log file
  	// also you can implement your own handlers, like write logs with http api
  	app.UseTimeoutHook(dotweb.DefaultTimeoutHookHandler, time.Second * 10)

  	app.HttpServer.GET("/index", func(ctx dotweb.Context) error{
  		return ctx.WriteString("welcome to my first web!")
  	})

  	//begin server
  	err := app.StartServer(80)
      fmt.Println("dotweb.StartServer error => ", err)
  }
  ~~~