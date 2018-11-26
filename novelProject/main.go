package main

import (
	"github.com/astaxie/beego"
	_ "novelProject/routers"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}

