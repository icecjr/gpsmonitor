package main

import (
	"fmt"
	. "project/elangshen/message"
	. "project/elangshen/mutils"
	"strings"
)

/* GPS定位包
 GPS: 日期 6byte
	  卫星数 1byte
	  纬度   4
	  经度   4
	  速度   1
	  航向   2
 LBS：MCC   2
	  MNC   1
	  LAC   2
	  CellID 3
 ACC    1
 数据上报模式 1
 GPS实时补传 1
里程统计     4

*/
// GT740 GV25 S17 G17
type package22 struct {
	g *GpsMessage
	l *LbsMessage
}

func (r *package22) dealPackage(deviceId string, messageStr []string) {
	r.g.DealPackage(deviceId, nil, messageStr)
	//计算数据库表的名字
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
}

// type gpsMessage struct {
// 	d_id            string    `bson:"d_id"`
// 	i_no            string    `bson:"i_no"`
// 	gps_wd          float64   `bson:"gps_wd"`
// 	gps_jd          float64   `bson:"gps_jd"`
// 	gps_speed       int       `bson:"gps_speed"`
// 	gps_real_dir    int       `bson:"gps_real_dir"`
// 	gps_dir_extends int       `bson:"gps_dir_extends"`
// 	d_time          time.Time `bson:"d_time"`
// 	c_time          time.Time `bson:"c_time"`
// 	d_version       int       `bson:"d_version"`
// 	gdgps_wd        float64   `bson:"gdgps_wd"`
// 	gdgps_jd        float64   `bson:"gdgps_jd"`
// }
// type gpsAndlbsMessage struct {
// 	D_id            string    `bson:"d_id"`
// 	I_no            string    `bson:"i_no"`
// 	Gps_wd          float64   `bson:"gps_wd"`
// 	Gps_jd          float64   `bson:"gps_jd"`
// 	Gps_speed       int       `bson:"gps_speed"`
// 	Gps_real_dir    int       `bson:"gps_real_dir"`
// 	Gps_dir_extends int       `bson:"gps_dir_extends"`
// 	D_time          time.Time `bson:"d_time"`
// 	C_time          time.Time `bson:"c_time"`
// 	D_version       int       `bson:"d_version"`
// 	Gdgps_wd        float64   `bson:"gdgps_wd"`
// 	Gdgps_jd        float64   `bson:"gdgps_jd"`
// }
// type lbsMessageStruct struct {
// 	d_id      string
// 	i_no      string
// 	lbs_mcc   string
// 	lbs_mnc   string
// 	lbs_lac   string
// 	cc_id     int
// 	c_sinal   int
// 	nearbts   string //一个序列
// 	wifi_info string
// 	d_time    time.Time
// 	c_time    time.Time
// 	d_version int
// }

/*
 GPS: 日期 6byte
	  卫星数 1byte
	  纬度   4
	  经度   4
	  速度   1
	  航向   2
 LBS：MCC   2
	  MNC   1
	  LAC   2
	  CellID 3
 ACC    1
 数据上报模式 1
 GPS实时补传 1
 序列号      2

*/
// G710 gps
// func (r *gpsMessage) dealPackage(deviceId string, moment []string, message []string) {
// 	r.d_id = deviceId
// 	r.d_time = TranferByteToTime(moment)
// 	r.i_no = "-1"
// 	r.d_version = 1
// 	r.gps_wd = TranferStrToLatitudeOrLogitude(message[1:5])
// 	r.gps_jd = TranferStrToLatitudeOrLogitude(message[5:9])
// 	r.gps_speed = StrToI(message[9])
// 	r.gps_real_dir = TranferStringToGpsDir(message[10:12])
// 	r.gps_dir_extends = 0

// 	r.gdgps_wd, r.gdgps_jd = ConvertGD(r.gps_wd, r.gps_jd)

// 	//计算数据库表的名字
// 	tableNo := fmt.Sprintf("%v", GetTableNo(deviceId))
// 	dataSheetName := strings.Join([]string{"gps_data_", tableNo}, "")
// 	//写入到数据库
// 	//fmt.Printf("deviceid %s datasheet name is %s", r.D_id, dataSheetName)
// 	//MongoDb.Insert(dataSheetName, bson.M{"name": "test", "Age": "34"})
// 	MongoDb.Insert(dataSheetName, r)

// }

// G740 0x19 8个字节 无时间标志
//lbs 9个字节 p17 LBS信息 struct 9
// LBS长度 1
// MCC 2
// MNC  1
// LAC 2
// Cell ID 3
// func (l *lbsMessage) dealPackage(deviceId string, moment []string, message []string) {
// 	l.d_id = deviceId
// 	if moment == nil {
// 		l.d_time = time.Now()
// 	} else {
// 		l.d_time = tranferByteToTime(moment)
// 	}

// 	l.i_no = "-1"
// 	l.lbs_mcc = string(strToI(message[0])*256 + strToI(message[1]))
// 	l.lbs_mnc = string(strToI(message[2]))
// 	l.lbs_lac = string(strToI(message[3])*256 + strToI(message[4]))
// 	l.cc_id = strToI(message[5])*256*256 + strToI(message[6])*256 + strToI(message[7])
// 	//TODO 需要继续处理
// 	l.c_sinal = 0
// 	// l.wifi_info = nil
// 	// l.nearbts = nil
// 	l.d_version = 1
// 	tableNo := fmt.Sprintf("%v", getTableNo(deviceId))
// 	dataSheetName := strings.Join([]string{"lbs_data_", tableNo}, "")
// 	//dataSheetName := "lbs_data_" + string(Crc16(l.d_id))
// 	MongoDb.Insert(dataSheetName, l)
// 	//更新 device_info l_time a_time wifi_info
// 	//转换成GPS信息
// }
// func (r *gpsAndlbsMessage) dealPackage(deviceId string, messageStr []string) {
// 	//TODO 把结构写到数据库中
// 	var lbsMessage lbsMessageStruct

// 	lbsMessage.d_id = deviceId
// 	r.D_id = deviceId
// 	tableNo := fmt.Sprintf("%v", GetTableNo(r.D_id))
// 	//fmt.Printf("d_id %s tableNo %s\n", r.D_id, TableNo)
// 	r.D_time = TranferByteToTime(messageStr[1:7])
// 	lbsMessage.d_time = r.D_time
// 	r.C_time = time.Now()
// 	lbsMessage.c_time = r.C_time
// 	buildGpsStruct(r, messageStr, tableNo)
// 	lbsMessage.d_id = deviceId
// 	buildLbsStruct(&lbsMessage, messageStr[19:], tableNo)
// 	MongoDb.Update("device_info", bson.M{"d_id": deviceId}, bson.M{"$set": bson.M{"a_time": r.C_time}})
// }

//TODO 8小时时差处理
// [00:00:05]:[1004][debug]cn.hhlcw.command.server.ClientServiceThread.java run:70行:deal logic data=78:78:22:22:12:04:0f:10:00:04:cf:03:4b:9b:58:0b:23:93:f0:6f:14:3a:01:cc:00:81:07:00:5b:7e:00:00:00:02:5a:66:40:0d:0a
// [00:00:05]:[1004][debug]cn.hhlcw.command.docking.GT710_Request.java request:53行:begin invoke method = do_gps_only
// [00:00:05]:[1004][debug]cn.hhlcw.extend.client.request.GT710_Device_Request.java do_gps_only:260行:{"data_device_time":1523808004000,"data_version":0,"device_id":"868120170641689","gps_dir_extends":5,"gps_jd":103.822782,"gps_real_dir":58,"gps_speed":111,"gps_status":0,"gps_wd":30.714787}
// func buildGpsStruct(r *gpsAndlbsMessage, messageStr []string, tableNo string) {
// 	r.I_no = "-1"
// 	r.D_version = 1

// 	r.Gps_wd = TranferStrToLatitudeOrLogitude(messageStr[8:12])
// 	r.Gps_jd = TranferStrToLatitudeOrLogitude(messageStr[12:16])
// 	r.Gps_speed = StrToI(messageStr[16])
// 	r.Gps_real_dir = TranferStringToGpsDir(messageStr[17:19])
// 	r.Gps_dir_extends = 0

// 	r.Gdgps_wd, r.Gdgps_jd = ConvertGD(r.Gps_wd, r.Gps_jd)

// 	//计算数据库表的名字

// 	dataSheetName := strings.Join([]string{"gps_data_", tableNo}, "")
// 	//写入到数据库
// 	//fmt.Printf("deviceid %s datasheet name is %s", r.D_id, dataSheetName)
// 	//MongoDb.Insert(dataSheetName, bson.M{"name": "test", "Age": "34"})
// 	MongoDb.Insert(dataSheetName, r)

// }

// func buildLbsStruct(l *lbsMessageStruct, messageStr []string, tableNo string) {
// 	//写入到数据库
// 	l.i_no = "-1"
// 	l.lbs_mcc = string(StrToI(messageStr[0])*256 + StrToI(messageStr[1]))
// 	l.lbs_mnc = string(StrToI(messageStr[2]))
// 	l.lbs_lac = string(StrToI(messageStr[3])*256 + StrToI(messageStr[4]))
// 	l.cc_id = StrToI(messageStr[5])*256*256 + StrToI(messageStr[6])*256 + StrToI(messageStr[7])
// 	l.c_sinal = 0
// 	// l.wifi_info = nil
// 	// l.nearbts = nil
// 	l.d_version = 1
// 	dataSheetName := strings.Join([]string{"lbs_data_", tableNo}, "")
// 	//dataSheetName := "lbs_data_" + string(Crc16(l.d_id))
// 	MongoDb.Insert(dataSheetName, l)

// }
