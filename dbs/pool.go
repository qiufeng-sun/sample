package dbs

import "util/dbs/redis"

var g_samplePools *redis.RedisPools

//
func getSampleConn(uid string) *redis.RedisConn {
	return g_samplePools.GetConn()
}

func Init() {
	redis.Init()

	g_samplePools = redis.GetRedisPools("sample")
}

//
func HealthCheck() error {
	return redis.HealthCheck()
}
