package routers

import (
	"techblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category/?:action", &controllers.CategoryController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.AutoRouter(&controllers.ArticleController{})
}
