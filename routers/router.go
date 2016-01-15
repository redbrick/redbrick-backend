package routers

import (
	"github.com/redbrick/redbrick-backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
