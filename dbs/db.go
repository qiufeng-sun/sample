package dbs

//import (
//	"fmt"

//	"mining/res"
//)

//func instanceId() int {
//	return res.GetMiningState().InstanceId
//}

//func userKey(userId int64) string {
//	//return "u" + userId + "_" + instanceId()
//	return fmt.Sprintf("u%v_%v", userId, instanceId())
//}

//////////////////////////////////////////////////////////////

////
//func GetDBMining(userId int64, zoneId int, conn *redis.RedisConn) (d *datas.DBMining, e *errs.Error) {
//    key := userKey(userId)
//	reply, err := ghRedis.String(conn.Do("GET", key))
//	if nil == err {
//		return datas.ParseDBMining(reply)
//	}

//	if ghRedis.ErrNil == err {
//		return nil, nil
//	}

//	return nil, errs.New(errs.MINING_GOT_FAILED, "get DBMining failed! userId=%v, zoneId=%v, err=%v", userId, zoneId, err.Error())
//}

//func SetDBMining(dbData *datas.DBMining, conn *redis.RedisConn) *errs.Error {
//    key := userKey(dbData.UserId)
//    val, _ := datas.ToDBString(dbData)

//	_, err := conn.Do("SET", key, val)
//	if err != nil {
//		return errs.New(errs.MINING_SAVE_FAILED, "dbs.SetDBMining() failed! userId=%v, zoneId=%v, reason=%v", dbData.UserId, dbData.ZoneId, err.Error())
//	}

//	return nil
//}
