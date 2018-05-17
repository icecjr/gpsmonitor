package message

import (
	"time"

	. "project/elangshen/mutils"
)

/*
GPS: 日期 6byte
卫星数 1byte
纬度   4
经度   4
速度   1
航向   2
*/
type GpsMessage struct {
	d_id            string    `bson:"d_id"`
	i_no            string    `bson:"i_no"`
	gps_wd          float64   `bson:"gps_wd"`
	gps_jd          float64   `bson:"gps_jd"`
	gps_speed       int       `bson:"gps_speed"`
	gps_real_dir    int       `bson:"gps_real_dir"`
	gps_dir_extends int       `bson:"gps_dir_extends"`
	d_time          time.Time `bson:"d_time"`
	c_time          time.Time `bson:"c_time"`
	d_version       int       `bson:"d_version"`
	gdgps_wd        float64   `bson:"gdgps_wd"`
	gdgps_jd        float64   `bson:"gdgps_jd"`
}

func (r *GpsMessage) DealPackage(deviceId string, moment []string, message []string) {
	r.d_id = deviceId
	r.d_time = TranferByteToTime(moment)
	r.i_no = "-1"
	r.d_version = 1
	r.gps_wd = TranferStrToLatitudeOrLogitude(message[1:5])
	r.gps_jd = TranferStrToLatitudeOrLogitude(message[5:9])
	r.gps_speed = StrToI(message[9])
	r.gps_real_dir = TranferStringToGpsDir(message[10:12])
	r.gps_dir_extends = 0

	r.gdgps_wd, r.gdgps_jd = ConvertGD(r.gps_wd, r.gps_jd)

	//计算数据库表的名字
	// tableNo := fmt.Sprintf("%v", GetTableNo(deviceId))
	// dataSheetName := strings.Join([]string{"gps_data_", tableNo}, "")
	// fmt.Println(dataSheetName)
	//写入到数据库
	//fmt.Printf("deviceid %s datasheet name is %s", r.D_id, dataSheetName)
	//MongoDb.Insert(dataSheetName, bson.M{"name": "test", "Age": "34"})
	//MongoDb.Insert(dataSheetName, r)

}
