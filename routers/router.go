// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api-demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/app",
			beego.NSInclude(
				&controllers.AppController{},
			),
		),
	)
	beego.AddNamespace(ns)
<<<<<<< HEAD
	// 设置路由
	beego.Router("/:type", &controllers.HttplibController{}, "get:Get;post:Post")
=======
>>>>>>> da5c6131c4b20775664dc9fee76bc60f3b1638d2
}
