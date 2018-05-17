/*
 */
package main

import (
	. "project/elangshen/message"
	//. "project/elangshen/mutils"
)

/* GT740 LBS报警包
报文头 4字节
报文内容 13字节
   MCC 2字节
   MNC 1字节
   LAC 2字节
   Cell ID 3字节
   终端信息 1
   电压信息  1
   GSM信号等级 1
   报警、语言 2

*/
type package19 struct {
	l *LbsMessage
	w *WarningMessage
}

func (r *package19) dealPackage(deviceId string, messageStr []string) {
	r.l.DealPackage(deviceId, nil, messageStr[4:12])
	r.w.DealPackage(deviceId, nil, messageStr[12:17])
}

// type lbsAndMultistationMessage struct {
// 	l *LbsMessage
// 	p *[6]wifiMessage
// }
// type wifiMessage struct {
// 	sig int
// 	id  int
// }

// //78:78:3B:28:12:05:06:07:13:14:01:CC:00:33:0F:00:2F:5A:3B:33:0F:00:9F:77:32:33:0F:00:9F:75:24:33:0F:00:9E:EA:1E:33:0F:00:9E:E0:1D:00:00:00:00:00:00:00:00:00:00:00:00:FF:00:01:00:00:A4:90:0D:0A
// func (r *lbsAndMultistationMessage) dealPackage(deviceId string, messageStr []string) {
// 	//TODO 处理
// 	//tableNo := fmt.Sprintf("%v", getTableNo(deviceId))
// 	//TODO 取出时间
// 	//moment := tranferByteToTime(messageStr[1:7])
// 	//r.l.DealPackage(deviceId, messageStr[1:7], messageStr[7:16])
// 	//buildLbsStruct(r.l, messageStr[1:], tableNo)
// 	//TODO 电压 GSM信息 更新 device_info
// 	for i := 0; i < 6; i++ {
// 		start := 16 + 6*i
// 		r.p[i].dealPackage(deviceId, messageStr[start:start+6])
// 	}
// 	//TODO写数据库 这个写库要麻烦一点

// }
// func (r *wifiMessage) dealPackage(deviceId string, messageStr []string) {
// 	r.sig = StrToI(messageStr[5])
// 	r.id = StrToI(messageStr[2])*256*256 + StrToI(messageStr[3])*256 + StrToI(messageStr[4])
// }
