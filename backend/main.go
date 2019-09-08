package main

import (
	"flux-web/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	h := controllers.InitHub()
	go h.Run()

	apiNs := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSRouter("/namespaces", &controllers.NamespaceController{}, "get:ListNamespaces"),
			beego.NSRouter("/workloads/:ns", &controllers.WorkloadController{}, "get:ListWorkloads"),
			beego.NSGet("/health", func(ctx *context.Context) {
            	ctx.Output.Body([]byte("All good"))
        	}),
		),
	)

	release := beego.NewNamespace("/api",
		beego.NSNamespace("/v1/release",
			beego.NSBefore(controllers.Auth),
            beego.NSRouter("/", &controllers.WorkloadController{}, "post:ReleaseWorkloads"),
		),
	)

	beego.AddNamespace(apiNs)
	beego.AddNamespace(release)

	beego.Router("/ws/v1", &controllers.WebSocketController{}, "get:ReleaseWorkloads")

	beego.Run()
}

