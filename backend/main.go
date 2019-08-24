package main

import (
	"flux-web/controllers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/api/v1", &controllers.WorkloadController{}, "get:ListWorkloads")
	//beego.Router("/namespaces", &controllers.WorkloadController{}, "get:ListServices")
	//beego.Router("/workloads/:ns", &controllers.WorkloadController{}, "get:ListWorkloads")
	//beego.Router("/release", &controllers.WorkloadController{}, "post:ReleaseWorkloads")

	beego.Router("/ws/v1", &controllers.WebSocketController{}, "get:ReleaseWorkloads")

	beego.Run()
}

