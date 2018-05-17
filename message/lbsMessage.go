package message

import (
	. "project/elangshen/mutils"
	"time"
)

/*
 LBS：MCC   2
	  MNC   1
	  LAC   2
	  CellID 3
*/
type lbs_cell struct {
	sig int
	id  int
}
type wifi_struct struct {
	mac string
	sig int
}
type LbsMessage struct {
	d_id        string
	i_no        string
	lbs_mcc     string
	lbs_mnc     string
	lbs_lac     string
	cc_id       int
	c_signal    int
	lbs_cell_id [6]lbs_cell
	wifi_info   []wifi_struct
	d_time      time.Time
	c_time      time.Time
	d_version   int
}

//
//lbs 9个字节 p17 LBS信息 struct 9
// LBS长度 1
// MCC 2
// MNC  1
// LAC 2
// Cell ID 3
func (l *LbsMessage) DealPackage(deviceId string, moment []string, message []string) {
	l.d_id = deviceId
	if moment == nil {
		l.d_time = time.Now()
	} else {
		l.d_time = TranferByteToTime(moment)
	}

	l.i_no = "-1"
	l.lbs_mcc = string(StrToI(message[0])*256 + StrToI(message[1]))
	l.lbs_mnc = string(StrToI(message[2]))
	l.lbs_lac = string(StrToI(message[3])*256 + StrToI(message[4]))
	l.cc_id = StrToI(message[5])*256*256 + StrToI(message[6])*256 + StrToI(message[7])
	//TODO 需要继续处理
	l.c_signal = 0
	// l.wifi_info = nil
	// l.nearbts = nil
	l.d_version = 1
	//TODO 区分基站
	if len(message) > 8 {
		for i := 0; i < 6; i++ {
			start := 16 + 6*i
			l.lbs_cell_id[i].dealPackage(message[start : start+6])
		}
	}
	//TODO 如果还有信息则是WIFI信息
	if len(message) > 44 {
		//TODO 读取wifi数量，循环处理wifi信息
	}
}
func (r *lbs_cell) dealPackage(messageStr []string) {
	r.sig = StrToI(messageStr[5])
	r.id = StrToI(messageStr[2])*256*256 + StrToI(messageStr[3])*256 + StrToI(messageStr[4])
}
func (r *wifi_struct) dealPackage(messageStr []string) {
	r.sig = StrToI(messageStr[5])
	r.mac = "xxx"
}
