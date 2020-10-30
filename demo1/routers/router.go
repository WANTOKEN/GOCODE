package routers

import (
	"demo1/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

    beego.Router("/", &controllers.MainController{})

    //基础路由
	beego.Get("/get", func(context *context.Context){
		beego.Info("基础路由中的get请求")
		context.Output.Body([]byte("基础路由中的get请求 get method"))
	})

	beego.Post("/post", func(context *context.Context){
		beego.Info("基础路由中的post请求")
		context.Output.Body([]byte("基础路由中的post请求 post method"))
	})

    //固定路由
	beego.Router("/GetInfo", &controllers.MainController{})

    //正则路由
    //*全匹配
	//beego.Router("/*", &controllers.MainController{})

	//:id变量匹配
	//beego.Router("/getUser/:id", &controllers.MainController{})

	//*.*匹配
	beego.Router("/upload/*.*", &controllers.MainController{})

    //自定义路由
	beego.Router("/getUserInfo", &controllers.MainController{},"GET:GetUserInfo")

	beego.Router("/getUser", &controllers.MainController{},"GET:GetUser")
}
