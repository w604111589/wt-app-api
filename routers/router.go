// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"wt-app-api/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "content_type", "token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true}))
	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/object",
	// 		beego.NSInclude(
	// 			&controllers.ObjectController{},
	// 			&controllers.ArticleController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&controllers.UserController{},
	// 		),
	// 	),
	// 	// beego.NSNamespace("/article",
	// 	// 	beego.NSInclude(
	// 	// 		&controllers.ArticleController{},
	// 	// 	),
	// 	// ),
	// )
	// beego.AddNamespace(ns)
	beego.AutoRouter(&controllers.ArticleController{})
	beego.AutoRouter(&controllers.LabelController{})
	beego.AutoRouter(&controllers.GoodsController{})
}
