package main

import (
	"runtime/debug"

	"util/logs"

	"sample/controllers"
	"sample/dbs"
	"sample/res"
)

// sample
func main() {
	defer func() {
		if e := recover(); e != nil {
			logs.GetLogger().Critical("panic:%v", e)
			logs.Warn(string(debug.Stack()))
		}
		logs.Close()
	}()

	initServer()

	controllers.Serve()
}

func initServer() {
	res.Init()

	dbs.Init()
}
