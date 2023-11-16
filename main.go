package main

import (
 "github.com/beego/beego/v2/server/web"
 "github.com/beego/beego/v2/server/web/context"
)

func main() {
 web.Get("/", func(ctx *context.Context) {
  ctx.WriteString("Hello World")
 })

 web.BConfig.Listen.HTTPAddr = "0.0.0.0"
 web.BConfig.Listen.HTTPPort = 8080

 web.Run()
}
