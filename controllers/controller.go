package controllers

import (
	"github.com/astaxie/beego"

	"sample/bizs"
)

///////////////////////////////////////////////////////////////////////////////
type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	//
	id := this.GetString("uid")

	this.Data["json"] = beego.BConfig.AppName + ":test -- uid=" + id
	this.ServeJSON()
}

///////////////////////////////////////////////////////////////////////////////
type HealthCheckController struct {
	beego.Controller
}

func (this *HealthCheckController) Get() {
	result := bizs.HandleHealthCheck()
	this.Ctx.WriteString(result)
}
