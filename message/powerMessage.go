package message

import (
	. "project/elangshen/mutils"
	"time"
)

/*
GT740 6位
  终端信息 1
  电压等级 2
  GSM信号强度 1
  语言/扩展口状态 2

  GV25 0x13
报文头 4字节
电压GSM内容 5字节 (电压是一个字节)
*/
type PowerMessage struct {
	D_id        string
	Gps_power_p int
	Power_info  string
	D_time      time.Time
	C_time      time.Time
	D_version   int
}

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

// deal logic data=78:78:0a:13:04:06:04:00:01:02:4d:2e:6e:0d:0a
// begin invoke method = do_status_13
// device_info=4,device_pow=6,gsm_signed=4,lang_ext=1
// power={"d_id":"868120170636358","d_time":1523808000148,"d_version":0,"gps_power_p":96}
func (r *PowerMessage) DealPackage(deviceId string, messageStr []string) {
	//把结构写到数据库中
	// 协议 0x13中电压是一个字节分为7级

	r.D_id = deviceId
	r.D_time = time.Now()
	// r.power_info = messageStr[1]
	r.Gps_power_p = handle_power(messageStr)
	r.D_version = 1
	//如果一个小时记录过一次，则不再记录
	//MongoDb.Insert("a_power_info", r)
	//更新 device_info表，power_time 为r.d_time d_power字段为r.gps_power_p
	//MongoDb.Update("device_info", bson.M{"d_id": deviceId}, bson.M{"$set": bson.M{"a_time": r.d_time}})

}

func handle_power(messageStr []string) int {
	var a [55]int = [55]int{0, 2, 5, 7, 10, 12, 15, 17, 20, 25, 30, 35, 40, 45, 48, 50, 52, 54, 55, 57, 59, 60, 62, 64, 68, 70, 71, 72, 73, 74, 75, 77, 78, 79, 80, 81, 83, 84, 85, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	if messageStr[0] == "13" {
		return StrToI(messageStr[2]) * 16
	} else if messageStr[0] == "23" {
		votage := StrToI(messageStr[2])*256 + StrToI(messageStr[3])
		if votage >= 365 && votage <= 420 {
			return a[votage-365]
		} else if votage > 420 {
			return 100
		}
	}
	return 0
}
