package message

import (
	"time"
)

type StatusMessage struct {
	d_id        string
	gps_power_p int
	power_info  string
	d_time      time.Time
	c_time      time.Time
	d_version   int
}

func (r *StatusMessage) DealPackage(deviceId string, messageStr []string) {
	//把结构写到数据库中
	// 协议 0x13中电压是一个字节分为7级

	r.d_id = deviceId
	r.d_time = time.Now()
	// r.power_info = messageStr[1]
	r.gps_power_p = handle_power(messageStr)
	r.d_version = 1
	//如果一个小时记录过一次，则不再记录
	//MongoDb.Insert("a_power_info", r)
	//更新 device_info表，power_time 为r.d_time d_power字段为r.gps_power_p
	//	MongoDb.Update("device_info", bson.M{"d_id": deviceId}, bson.M{"$set": bson.M{"a_time": r.d_time}})

}
