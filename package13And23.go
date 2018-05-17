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
/* 0x23包结构 GT740 S17 心跳包
报文头 4字节
电压GSM内容 6字节
报文尾     6字节 （电压是两个字节）
GV25 0x13
报文头 4字节
电压GSM内容 5字节 (电压是一个字节)
报文尾     6字节
*/
type package13And23 struct {
	l *PowerMessage
}

func (r *package13And23) dealPackage(deviceId string, messageStr []string) {
	//把结构写到数据库中
	// 协议 0x13中电压是一个字节分为7级
	r.l.DealPackage(deviceId, messageStr)
	//如果一个小时记录过一次，则不再记录
	MongoDb.Insert("a_power_info", r.l)
	//更新 device_info表，power_time 为r.d_time d_power字段为r.gps_power_p
	MongoDb.Update("device_info", bson.M{"d_id": deviceId}, bson.M{"$set": bson.M{"a_time": r.l.D_time}})
}

// type devicePowMessage struct {
// 	d_id        string
// 	gps_power_p int
// 	power_info  string
// 	d_time      time.Time
// 	c_time      time.Time
// 	d_version   int
// }
// type statusMessage struct {
// 	d_id        string
// 	gps_power_p int
// 	power_info  string
// 	d_time      time.Time
// 	c_time      time.Time
// 	d_version   int
// }

// func (r *statusMessage) dealPackage(deviceId string, messageStr []string) {
// 	//把结构写到数据库中
// 	// 协议 0x13中电压是一个字节分为7级

// 	r.d_id = deviceId
// 	r.d_time = time.Now()
// 	// r.power_info = messageStr[1]
// 	r.gps_power_p = handle_power(messageStr)
// 	r.d_version = 1
// 	//如果一个小时记录过一次，则不再记录
// 	MongoDb.Insert("a_power_info", r)
// 	//更新 device_info表，power_time 为r.d_time d_power字段为r.gps_power_p
// 	MongoDb.Update("device_info", bson.M{"d_id": deviceId}, bson.M{"$set": bson.M{"a_time": r.d_time}})

// }

// // deal logic data=78:78:0a:13:04:06:04:00:01:02:4d:2e:6e:0d:0a
// // begin invoke method = do_status_13
// // device_info=4,device_pow=6,gsm_signed=4,lang_ext=1
// // power={"d_id":"868120170636358","d_time":1523808000148,"d_version":0,"gps_power_p":96}
// func (r *devicePowMessage) dealPackage(deviceId string, messageStr []string) {
// 	//把结构写到数据库中
// 	// 协议 0x13中电压是一个字节分为7级

// 	r.d_id = deviceId
// 	r.d_time = time.Now()
// 	// r.power_info = messageStr[1]
// 	r.gps_power_p = handle_power(messageStr)
// 	r.d_version = 1
// 	//如果一个小时记录过一次，则不再记录
// 	MongoDb.Insert("a_power_info", r)
// 	//更新 device_info表，power_time 为r.d_time d_power字段为r.gps_power_p
// 	MongoDb.Update("device_info", bson.M{"d_id": deviceId}, bson.M{"$set": bson.M{"a_time": r.d_time}})

// }

// func handle_power(messageStr []string) int {
// 	var a [55]int = [55]int{0, 2, 5, 7, 10, 12, 15, 17, 20, 25, 30, 35, 40, 45, 48, 50, 52, 54, 55, 57, 59, 60, 62, 64, 68, 70, 71, 72, 73, 74, 75, 77, 78, 79, 80, 81, 83, 84, 85, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
// 	if messageStr[0] == "13" {
// 		return StrToI(messageStr[2]) * 16
// 	} else if messageStr[0] == "23" {
// 		votage := StrToI(messageStr[2])*256 + StrToI(messageStr[3])
// 		if votage >= 365 && votage <= 420 {
// 			return a[votage-365]
// 		} else if votage > 420 {
// 			return 100
// 		}
// 	}
// 	return 0
// }
