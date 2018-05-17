/*
 */
package main

import (
	. "project/elangshen/message"
)

/* GT740 0x28 多基站扩展信息包
lbs信息
6个基站信息
*/
type package28 struct {
	l *LbsMessage
}

func (r *package28) dealPackage(deviceId string, messageStr []string) {
	r.l.DealPackage(deviceId, nil, messageStr)
}

// type lbsAndStatusMessage struct {
// 	l *LbsMessage
// 	p *[6]wifiMessage
// }

// func (r *lbsAndStatusMessage) dealPackage(deviceId string, messageStr []string) {
// 	//TODO 处理 第5个字段开始共8个是LBS信息，这里没有时间信息
// 	//	r.l.dealPackage(deviceId, nil, messageStr[4:12])
// 	//TODO 从13个字段开始 共5个字段是状态信息
// 	//	r.s.dealPackage(deviceId, messageStr[12:17])

// }
