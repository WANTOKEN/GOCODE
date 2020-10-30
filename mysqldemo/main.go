package main

import (
	_ "mysqldemo/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

