package controllers

import (
	"os"
	"strings"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["namespaces"] = strings.Split(os.Getenv("NAMESPACES"), ";")
	this.Data["fluxUrl"] = os.Getenv("FLUX_URL")
	this.Data["workloads"] = GetImages()
	this.TplName = "main.tpl"
}
