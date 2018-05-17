package message

import (
	"time"
	//. "project/elangshen/mutils"
)

/*
GPS: 日期 6byte
卫星数 1byte
纬度   4
经度   4
速度   1
航向   2
*/
type ResponseMessage struct {
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

func (r *ResponseMessage) DealPackage(deviceId string, moment []string, message []string) {
}
