package main

import (
	. "project/elangshen/message"
)

/* GT740 0x28 多WIFI扩展信息包
lbs信息
6个基站信息包
6个WIFI信息
*/
type package2c struct {
	l *LbsMessage
}

func (r *package2c) dealPackage(deviceId string, messageStr []string) {
	r.l.DealPackage(deviceId, nil, messageStr)
}
