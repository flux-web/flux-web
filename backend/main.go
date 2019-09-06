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

	releaseNs := beego.NewNamespace("/api",
		beego.NSNamespace("/rel/v1",
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

	apiNs := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSRouter("/namespaces", &controllers.NamespaceController{}, "get:ListNamespaces"),
			beego.NSRouter("/workloads/:ns", &controllers.WorkloadController{}, "get:ListWorkloads"),
			beego.NSGet("/health", func(ctx *context.Context) {
            	ctx.Output.Body([]byte("All good"))
        	}),
		),
	)

	beego.AddNamespace(releaseNs)
	beego.AddNamespace(apiNs)

	beego.Router("/ws/v1", &controllers.WebSocketController{}, "get:ReleaseWorkloads")

	beego.Run()
}

