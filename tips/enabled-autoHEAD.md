* How To:
  - how to add support "HEAD" method for all requests in your httpserver?
  ~~~go
  func main() {
    app := dotweb.Classic("/home/logs/wwwroot/")

    // if use this, all router will auto add "HEAD" method support
    // default is false
    app.HttpServer.SetEnabledAutoHEAD(true)

    app.HttpServer.GET("/index", func(ctx dotweb.Context) error{
        return ctx.WriteString("welcome to my first web!")
    })

    err := app.StartServer(80)
    fmt.Println("dotweb.StartServer error => ", err)
  }
  ~~~