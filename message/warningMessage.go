package message

import (
	. "project/elangshen/mutils"
	"time"
)

/*
   终端信息 1
   电压信息  1
   GSM信号等级 1
   报警、语言 2
*/
var warning_map map[string]string = map[string]string{
	"01": "SOS求救",
	"02": "断电报警",
	"03": "震动报警",
	"04": "进围栏报警",
	"05": "出围栏报警",
	"06": "超速报警",
	"09": "位移报警",
	"0A": "进 GPS 盲区报警",
	"0B": "出 GPS 盲区报警",
	"0C": "开机报警",
	"0D": "GPS 第一次定位报警",
	"0E": "外电低电报警",
	"0F": "外电低电保护报警",
	"10": "换卡报警",
	"11": "关机报警",
	"12": "外电低电保护后飞行模式报警",
	"13": "拆卸报警",
	"14": "门报警",
	"15": "低电关机报警",
	"16": "声控报警",
	"17": "伪基站报警",
	"FF": "ACC 关",
}

type position struct {
	gps_jd      float64
	gps_wd      float64
	last_d_time time.Time
	last_pos    string
}
type power struct {
	power_level int
	gsm_sign    int
}
type WarningMessage struct {
	r_no      int //记录编号
	d_id      string
	org_id    string //门店ID
	bs_no     string //设备识别码
	d_code    string //报警编号,如果是围栏报警，1表示进围栏报警，2表示出围栏报警
	w_code    int    //报警类型 1: 设备报警，2：围栏报警, 3,:离线超时报警, 4:停留超时报警
	w_type    int
	w_note    string //报警描述
	w_extend  string
	p_extend  position
	x_extend  power
	d_time    time.Time //设备时间
	c_time    time.Time //记录添加时间
	d_version int
}

func (r *WarningMessage) DealPackage(deviceId string, moment []string, messageStr []string) {
	// 判断 w_code!=0，则处理
	r.w_code = StrToI(messageStr[3])
	if r.w_code == 0 {
		return
	}
	r.d_id = deviceId

	r.d_code = "8196"
	aPower := new(power)
	aPower.power_level = StrToI(messageStr[1])
	aPower.gsm_sign = StrToI(messageStr[2])
	r.w_code = StrToI(messageStr[3])
	r.w_type = 1
	r.w_note = warning_map[messageStr[3]]
	r.d_time = TranferByteToTime(moment)
	r.c_time = time.Now()
	r.d_version = 1

	//TODO 设置org_id 判断数据库是否已有该设备的信息
	//判断 device_info这个数据库中 d_time和l_time的值，这个两个值分别记录了Gps最后的记录时间以及lbs的记录时间
	//读出gps gps_wd gps_jd 或者lbs_wd lbs_jd放入r中的p_extend
	//var l deviceWarningMessage
	//	MongoDb.Find("device_info", bson.M{"d_id": deviceId}, &l)

	//更新数据库 a_warning
}
