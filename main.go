package main

import (
 "github.com/beego/beego/v2/server/web"
 "github.com/beego/beego/v2/server/web/context"
)

func main() {
 web.Get("/", func(ctx *context.Context) {
  ctx.WriteString("Hello World")
 })
 web.Run()
}
