/*
 */
package main

import (
	. "project/elangshen/message"
	//. "project/elangshen/mutils"
)

/*0x21 GT740 GV25 S17 G17
终端回复指令
5个字节报文头 0x79 0x79
服务器标志位 4
内容编码 1
内容 M 需要发送的数据
*/
//TODO 需要确认message的支持
type package21 struct {
	l *ResponseMessage
}

func (r *package21) dealPackage(deviceId string, messageStr []string) {
	r.l.DealPackage(deviceId, nil, messageStr)
	MongoDb.Insert("a_power_info", r.l)
}

// type commondSetResponseMessage struct {
// 	d_id        string
// 	gps_power_p int
// 	power_info  string
// 	d_time      time.Time
// 	c_time      time.Time
// 	d_version   int
// }

// func (r *commondSetResponseMessage) dealPackage(deviceId string, messageStr []string) {
// 	//TODO 处理

// 	r.d_id = deviceId
// 	r.d_time = time.Now()
// 	r.power_info = messageStr[1]
// 	r.gps_power_p = StrToI(messageStr[2]) * 16
// 	r.d_version = 1
// 	MongoDb.Insert("a_power_info", r)
// }
