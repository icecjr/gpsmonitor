package main

import (
	. "project/elangshen/message"

	"gopkg.in/mgo.v2/bson"
)

type package01 struct {
	l *LoginMessage
}

/*登陆信息包
起始位 2
包长度 1
协议号 1 //0x01
终端 ID 8
类型识别码 2 //两个字节前三位代表机种，后一位代表分支。如GT700，类型识别码是0x70 0x00
时区语言 2  //最低一位是8则西时区，0为东时区，然后右移四位后然后除以100得到时区
信息序列号 2
错误校验 2
停止位 2 //0x0D 0x0A
*/
func (r *package01) dealPackage(deviceId string, message []string) {

	//TODO 需要回复
	r.l.DealPackage(deviceId, message[4:16])
	//更新设备列表
	//MongoDb.Update("device_info", bson.M{"d_id": deviceId}, bson.M{"$set": bson.M{"a_time": r.a_time}})
	MongoDb.Update("device_info", bson.M{"d_id": r.l.DeviceId}, bson.M{"$set": bson.M{"a_time": r.l.A_time}})
}
