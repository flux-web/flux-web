package main

import (
	"github.com/astaxie/beego"
	"flux-web/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/workload/:ns", &controllers.WorkloadController{}, "get:ListWorkloads")
	beego.Router("/release", &controllers.WorkloadController{}, "post:ReleaseWorkloads")
	beego.Run()
}
