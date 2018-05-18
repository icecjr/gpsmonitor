package main

import (
	. "project/elangshen/message"
	//. "project/elangshen/mutils"

	"gopkg.in/mgo.v2/bson"
)

/*处理流程
如果一个小时内记录过一次，将不在记录
修改设备表中的电量信息
*/
/*
GV25 0x13
报文头 4字节
电压GSM内容 5字节 (电压是一个字节)
报文尾     6字节
*/
type package13 struct {
	l PowerMessage
}

func (r *package13) dealPackage(deviceId string, messageStr []string) {
	//把结构写到数据库中
	// 协议 0x13中电压是一个字节分为7级
	r.l.DealPackage(deviceId, messageStr)
	//TODO 如果一个小时记录过一次，则不再记录
	MongoDb.Insert("a_power_info", &r.l)
	//更新 device_info表，power_time 为r.d_time d_power字段为r.gps_power_p
	MongoDb.Update("device_info", bson.M{"d_id": deviceId}, bson.M{"$set": bson.M{"a_time": r.l.D_time}})
}
