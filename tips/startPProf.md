
* How To:
  - how to start golang pprof-server in your app?
  ~~~go
  func main() {
    app := dotweb.Classic("/home/logs/wwwroot/")


    // set enabled pprof http server on port 8081
    // you can curl http://127.0.0.1:8081/debug/pprof
    // don't use same port with StartServer
    // default is disable
    app.SetPProfConfig(true, 8081)

    app.HttpServer.GET("/index", func(ctx dotweb.Context) error{
        return ctx.WriteString("welcome to my first web!")
    })

    //begin server
    err := app.StartServer(80)
    fmt.Println("dotweb.StartServer error => ", err)
  }
  ~~~