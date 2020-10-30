package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

//固定路由post请求
func (this *MainController) Post() {
	beego.Info("固定路由post请求")
	this.Ctx.Output.Body([]byte("固定路由post请求"))
}

//固定路由Delete请求
func (this *MainController) Delete() {
	beego.Info("固定路由Delete请求")
	this.Ctx.Output.Body([]byte("固定路由Delete请求"))
}

//func (this *MainController) Get() {
// 	id := this.Ctx.Input.Param(":id")
// 	beego.Info("ID："+id)
// 	this.Ctx.ResponseWriter.Write([]byte("ID："+id))
//}

func (this *MainController) Get() {
	file := this.Ctx.Input.Param(":path")
	beego.Info("file："+file)
	ext := this.Ctx.Input.Param(":ext")
	beego.Info("ext："+ext)
	this.Ctx.ResponseWriter.Write([]byte("path："+file+",ext:"+ext))
}

func (this *MainController) GetUserInfo() {
	id := this.GetString("id")
	beego.Info("id："+id)
	name := this.GetString("name")
	beego.Info("name："+name)
	this.Ctx.ResponseWriter.Write([]byte("id："+id+",name:"+name))
}


func (this *MainController) GetUser() {
	this.Ctx.Output.Body([]byte("1"))
}
