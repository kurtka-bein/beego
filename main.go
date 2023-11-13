package main

import "github.com/beego/beego/v2/server/web"

func main() {
    web.Get("/", func(ctx *web.Context) {
        ctx.WriteString("Hello World")
    })
    web.Run()
}
