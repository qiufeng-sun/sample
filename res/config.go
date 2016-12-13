package res

import (
	//	"flag"

	"github.com/astaxie/beego"

	"util"
	"util/logs"
	"util/logs/scribe"
)

type AppConfig struct {
	ZkEnv string
}

var g_appConfig = &AppConfig{}

func GetConfig() *AppConfig {
	return g_appConfig
}

func Init() {
	//
	//	parseFlag()
}

////
//func parseFlag() {
//	ins := flag.String("ins", "get", "ins:[add, get]")
//	flag.Parse()

//	g_appConfig.Instance = *ins
//}

////////////////////////////////////////////////////////////////////////////////
func init() {
	//
	logs.Init("conf/log.conf")

	// zk env
	g_appConfig.ZkEnv = beego.AppConfig.String("zk_env")

	// scribe
	//	confd, e := config.NewConfig("ini", confFile)
	//	scribe.Init("sample", confd)
	scribe.Init("sample", beego.AppConfig)

	// log
	logs.Info("init: %v\n", util.ToJsonString(*g_appConfig))
}
