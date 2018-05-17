package main

import (
	"fmt"
	. "project/elangshen/message"
	. "project/elangshen/mutils"
	"strings"
)

/*报警包 0x26 GT740 GV25 S17 G17 gps报警包
报文头
gps信息
lbs信息
电压/GSM信息/报警信息
*/

type package26 struct {
	g *GpsMessage
	l *LbsMessage
	p *PowerMessage
}

func (r *package26) dealPackage(deviceId string, messageStr []string) {
	r.g.DealPackage(deviceId, nil, messageStr)
	tableNo := fmt.Sprintf("%v", GetTableNo(deviceId))
	dataSheetName := strings.Join([]string{"gps_data_", tableNo}, "")
	//写入到数据库
	//fmt.Printf("deviceid %s datasheet name is %s", r.D_id, dataSheetName)
	//MongoDb.Insert(dataSheetName, bson.M{"name": "test", "Age": "34"})
	MongoDb.Insert(dataSheetName, r.g)
	r.l.DealPackage(deviceId, nil, messageStr)
	dataSheetName = strings.Join([]string{"lbs_data_", tableNo}, "")
	//dataSheetName := "lbs_data_" + string(Crc16(l.d_id))
	MongoDb.Insert(dataSheetName, r.l)
	r.p.DealPackage(deviceId, messageStr)
}

//TODO  电量映射表 确定我们设备的种类
// var warning_map map[string]string = map[string]string{
// 	"01": "SOS求救",
// 	"02": "断电报警",
// 	"03": "震动报警",
// 	"04": "进围栏报警",
// 	"05": "出围栏报警",
// 	"06": "超速报警",
// 	"09": "位移报警",
// 	"0A": "进 GPS 盲区报警",
// 	"0B": "出 GPS 盲区报警",
// 	"0C": "开机报警",
// 	"0D": "GPS 第一次定位报警",
// 	"0E": "外电低电报警",
// 	"0F": "外电低电保护报警",
// 	"10": "换卡报警",
// 	"11": "关机报警",
// 	"12": "外电低电保护后飞行模式报警",
// 	"13": "拆卸报警",
// 	"14": "门报警",
// 	"15": "低电关机报警",
// 	"16": "声控报警",
// 	"17": "伪基站报警",
// 	"FF": "ACC 关",
// }

// type position struct {
// 	gps_jd      float64
// 	gps_wd      float64
// 	last_d_time time.Time
// 	last_pos    string
// }
// type power struct {
// 	power_level int
// 	gsm_sign    int
// }
// type stationAndWifiMessage struct {
// 	r_no      int //记录编号
// 	d_id      string
// 	org_id    string //门店ID
// 	bs_no     string //设备识别码
// 	d_code    string //报警编号,如果是围栏报警，1表示进围栏报警，2表示出围栏报警
// 	w_code    int    //报警类型 1: 设备报警，2：围栏报警, 3,:离线超时报警, 4:停留超时报警
// 	w_type    int
// 	w_note    string //报警描述
// 	w_extend  string
// 	p_extend  position
// 	x_extend  power
// 	d_time    time.Time //设备时间
// 	c_time    time.Time //记录添加时间
// 	d_version int
// }

//78:78:25:26:12:05:09:06:24:17:CC:04:57:1B:38:0C:5E:F8:90:00:54:7F:09:01:CC:00:30:A3:00:6F:55:50:00:04:02:01:00:09:24:9E:0D:0A
/* GT750 0x26 报警包 (GPS、LBS、状态合并信息包) P16
起始位 2byte 78:78
包长度 1 25
协议号 1 26
日期时间 6 12:05:09:06:24:17
GPS 信息 struct 12
GPS信息卫星数 1 CC
纬度   4
经度  4
速度 1
航向、状态 2
LBS信息 struct 9
LBS长度 1
MCC 2
MNC  1
LAC 2
Cell ID 3
状态信息 struct 5
终端信息内容 1
电压等级 1
GSM信号强度 1
报警/预言/扩展口状态 2
struct end
围栏编号 1 //多围栏才有这个条目 也就是说多围栏报文长度有43个字节，而单围栏报文长度有42个字节
序列号 2
错误校验 2
结束位  2

*/
// func (r *deviceWarningMessage) dealPackage(deviceId string, messageStr []string) {
// 	//TODO 报文头报文解析
// 	//gpsPackage := new(GpsMessage)
// 	//gpsPackage.DealPackage(deviceId, messageStr[1:7], messageStr[10:22])
// 	//TODO Gps报文解析
// 	//TODO LBS报文解析
// 	//lbsPackage := new(LbsMessage)
// 	//lbsPackage.DealPackage(deviceId, messageStr[1:7], messageStr[24:31])
// 	//TODO 报警信息报警解析

// 	warningPackage := new(WarningMessage)
// 	warningPackage.DealPackage(deviceId, messageStr[1:7], messageStr[31:36])
// }

// type warningMessage struct {
// 	r_no      int //记录编号
// 	d_id      string
// 	org_id    string //门店ID
// 	bs_no     string //设备识别码
// 	d_code    string //报警编号,如果是围栏报警，1表示进围栏报警，2表示出围栏报警
// 	w_code    int    //报警类型 1: 设备报警，2：围栏报警, 3,:离线超时报警, 4:停留超时报警
// 	w_type    int
// 	w_note    string //报警描述
// 	w_extend  string
// 	p_extend  []position
// 	x_extend  []power
// 	d_time    time.Time //设备时间
// 	c_time    time.Time //记录添加时间
// 	d_version int
// }
// type deviceWarningMessage struct {
// 	w *warningMessage
// }

// func (r *warningMessage) dealPackage(deviceId string, moment []string, messageStr []string) {
// 	// 判断 w_code!=0，则处理
// 	r.w_code = StrToI(messageStr[3])
// 	if r.w_code == 0 {
// 		return
// 	}
// 	r.d_id = deviceId

// 	r.d_code = "8196"
// 	aPower := new(power)
// 	aPower.power_level = StrToI(messageStr[1])
// 	aPower.gsm_sign = StrToI(messageStr[2])
// 	r.w_code = StrToI(messageStr[3])
// 	r.w_type = 1
// 	r.w_note = warning_map[messageStr[3]]
// 	r.d_time = TranferByteToTime(moment)
// 	r.c_time = time.Now()
// 	r.d_version = 1

// 	//TODO 设置org_id 判断数据库是否已有该设备的信息
// 	//判断 device_info这个数据库中 d_time和l_time的值，这个两个值分别记录了Gps最后的记录时间以及lbs的记录时间
// 	//读出gps gps_wd gps_jd 或者lbs_wd lbs_jd放入r中的p_extend
// 	var l deviceWarningMessage
// 	MongoDb.Find("device_info", bson.M{"d_id": deviceId}, &l)

// 	//更新数据库 a_warning
// }
