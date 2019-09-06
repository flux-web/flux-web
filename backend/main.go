package main

import (
	"os"
	"strconv"
	"flux-web/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	h := controllers.InitHub()
	go h.Run()

	writeNs := beego.NewNamespace("/api",
		beego.NSNamespace("/v1/write",
			beego.NSCond(func(ctx *context.Context) bool {
				if readOnly, err := strconv.ParseBool(os.Getenv("READ_ONLY")); err != nil  {
					return true
				} else{
					return !readOnly
				}
			}),
			beego.NSRouter("/release", &controllers.WorkloadController{}, "post:ReleaseWorkloads"),
		),
	)

	readNs := beego.NewNamespace("/api",
		beego.NSNamespace("/v1/read",
			beego.NSRouter("/namespaces", &controllers.NamespaceController{}, "get:ListNamespaces"),
			beego.NSRouter("/workloads/:ns", &controllers.WorkloadController{}, "get:ListWorkloads"),
			beego.NSGet("/health", func(ctx *context.Context) {
            	ctx.Output.Body([]byte("All good"))
        	}),
		),
	)

	beego.AddNamespace(writeNs)
	beego.AddNamespace(readNs)

	beego.Router("/ws/v1", &controllers.WebSocketController{}, "get:ReleaseWorkloads")

	beego.Run()
}

