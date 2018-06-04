package routers

import (
	"PhoneSaleManage/PhoneSaleManage/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
