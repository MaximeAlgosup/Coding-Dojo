package routers

import (
	"coding-dojo/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello-world", &controllers.MainController{}, "get:HelloSitepoint")
}