package main

import (
	_ "coding-dojo/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

