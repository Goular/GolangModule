package routers

import (
	"WEB/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/testInput", &controllers.TestInputController{}, "get:Get;post:Post")
	beego.Router("/testInput2", &controllers.TestInputController{}, "get:Get2")
	beego.Router("/testLogin", &controllers.TestLoginController{}, "get:Login;post:Post")
	beego.Router("/testModel", &controllers.TestModelController{}, "get:Get;post:Post")
	beego.Router("/testModel2", &controllers.TestModelController{}, "get:Get2")
	beego.Router("/testView", &controllers.TestViewController{}, "get:Get")
	beego.Router("/testConfig", &controllers.TestConfigController{}, "get:Get")
	beego.Router("/testHttpLib", &controllers.TestHttpLibController{}, "get:Get")
	beego.Router("/testContext", &controllers.TestContextController{}, "get:Get")

}
