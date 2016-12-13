package controllers

import (
	"github.com/astaxie/beego"
)

///////////////////////////////////////////////////////////////////////////////
//
func Serve() {
	regHandler()

	beego.Run()
}

//
func regHandler() {
	beego.Router("/monitor/sample/test", &TestController{}, "get:Get")
	beego.Router("/monitor/sample/healthcheck", &HealthCheckController{}, "get:Get")
}
