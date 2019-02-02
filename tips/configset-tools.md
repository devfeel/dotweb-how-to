* How To:
  - how to use ConfigSet tools in your app?
  - first define config file named 'user-conf.xml'
  ```
  <?xml version="1.0" encoding="UTF-8"?>
  <config>
      <set key="set1" value="1" />
      <set key="set2" value="2" />
      <set key="set3" value="3" />
      <set key="set4" value="4" />
  </config>
  ```
  - include config file
  ```
  err := app.Config.IncludeConfigSet("/home/your/user-conf.xml", config.ConfigType_XML)
  ```
  - use it in your HttpContext
  ```
  value := ctx.ConfigSet().GetString("set1")
  ```