/*
Elangshen system data receiver and dealer

Edit by cjr 2018-4-17
*/
package main

import (
	"flag"
	. "project/elangshen/mutils"

	"github.com/golang/glog"
	ini "gopkg.in/ini.v1"
)

const iniFilePath = "conf.ini" //配置文件路径

var mongodbConfig struct {
	MongoDb_hostname string
	DBName           string
	DBUser           string
	DBPassword       string
}

//var dbconfig mongodbConfig //mongodb Config 文件

var mqttConfig struct {
	Mqtt_broker string
	Topic       string
	ClientID    string
}
var tcpConfig struct {
	Tcp_url string
}

//var mqconfig mqttConfig

//init在main之前执行
func init() {
	flag.Parse()
	flag.Set("log_dir", "./logs")        // 1
	flag.Set("MaxSize", "1024*1024*500") //log文件最大尺寸500M
	conf, err := ini.Load(iniFilePath)   //加载配置文件
	CheckError(err)

	// 读mongodb配置
	//dbconfig := new(mongodbConfig)
	conf.BlockMode = false
	err = conf.Section("Mongodb").MapTo(&mongodbConfig) //解析成结构体
	// log.Println(conf.Section("Mongodb").Key("MongoDb_hostname").String())
	CheckError(err)

	//mqconfig := new(mqttConfig)
	err = conf.Section("Mqtt").MapTo(&mqttConfig) //解析成结构体
	CheckError(err)

	//log.Println(mongodbConfig, mqttConfig)
	glog.Info(mongodbConfig, mqttConfig)
	glog.Flush()
}

// func ConnectMongoDb() {
// 	//prepare mongodb *************************************************
// 	DbSessivar DbSession Session
// var gpsDatabase mgo.Databaseon, err := mgo.Dial(mongodbConfig.MongoDb_hostname)
// 	CheckError(err)
// 	defer DbSession.Close()
// 	// Optional. Switch the session to a monotonic behavior.
// 	DbSession.SetMode(mgo.Monotonic, true)
// 	fmt.Println("begin to connect mongodb")
// 	gpsDatabase := DbSession.DB(mongodbConfig.DBName)
// 	gpsDatabase.C("gps_test").Insert(bson.M{"name": "test", "Age": "34"})
// 	fmt.Println("End of connect mongodb")
// 	// db.C("test").Insert(bson.M{"name": "3", "value": "5"})
// 	// log.Println(db.Name)

// }
var MongoDb *DB

func main() {

	defer glog.Flush()
	MongoDb = new(DB)
	MongoDb.ConnectMongoDb()
	ConnectMqtt()
	MongoDb.CloseMongoDb()
	//CreateTCPConnection() //后面如果需要设备和服务器直接相连，注销ConnectMqtt函数，使用当前函数
}
