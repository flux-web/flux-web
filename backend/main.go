package main

import (
	"flux-web/controllers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.WorkloadController{}, "get:ListWorkloads")
	beego.Router("/workloads/:ns", &controllers.WorkloadController{}, "get:ListWorkloads")
	beego.Router("/release", &controllers.WorkloadController{}, "post:ReleaseWorkloads")
	beego.Run()
}

