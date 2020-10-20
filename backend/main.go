package main

import (
	"github.com/flux-web/flux-web/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
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

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Request-Method","Content-Type"},
		ExposeHeaders:    []string{"Access-Control-Allow-Origin"},
	}))

	beego.Run()
}
