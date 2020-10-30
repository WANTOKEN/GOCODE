// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apidemo/MyGoApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/one",
			beego.NSInclude(
				&controllers.OneController{},
			),
		),

		beego.NSNamespace("/three",
			beego.NSInclude(
				&controllers.ThreeController{},
			),
		),

		beego.NSNamespace("/two",
			beego.NSInclude(
				&controllers.TwoController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
