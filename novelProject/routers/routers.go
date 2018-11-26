package routers

import (
	"github.com/astaxie/beego"
	"novelProject/controllers"
)

func init() {
	beego.Include(&controllers.MainController{})
}

