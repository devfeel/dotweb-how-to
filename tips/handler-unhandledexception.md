* How To:
  - how to create a handler for an unhandled exception in your app?
  ~~~go
  func main() {
    app := dotweb.Classic("/home/logs/wwwroot/")

    // if there is an unhandled exception, dotweb will use this to handle it.
    // at this handler, you can log error message or do everything what you want.
    // also, dotweb will log the error request info in "dotweb_server_ERROR_yyyy_MM_dd.log".
    app.SetExceptionHandle(func(ctx dotweb.Context, err error) {
        ctx.WriteString("oh, we are panic!", err.Error())
    })

    // we created an exception, no handle it.
    app.HttpServer.GET("/index", func(ctx dotweb.Context) error{
        i := 0
        b := 2 / i
        return ctx.WriteString(b)
    })

    err := app.StartServer(80)
    fmt.Println("dotweb.StartServer error => ", err)
  }
  ~~~
