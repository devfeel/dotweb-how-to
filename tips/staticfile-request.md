* How To:
  - how to add support static files request in your app?
  ~~~go
  func main() {
    app := dotweb.Classic("/home/wwwroot/logs/")

    // if use ServerFile, default http method is GET
    app.HttpServer.ServerFile("/static/*", "/home/wwwroot/static/")

    // if use RegisterServerFile, you can specify other HttpMethod, like "POST"
    app.HttpServer.RegisterServerFile(dotweb.RouteMethod_POST, "/staticR/*", "/home/wwwroot/static/")

    err := app.StartServer(80)
    fmt.Println("dotweb.StartServer error => ", err)
  }
  ~~~