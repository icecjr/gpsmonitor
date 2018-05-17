package prepareKnow

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var filepath = "./test.ini"

type Config struct {
	MongoDb_hostname string
	DBName           string
	DBUser           string
	DBPassword       string
	// MongoDbHostname string `ini:"MongoDb_hostname"`
	// DBName          string `ini:"DBname"`
	// DBUser          string `ini:"DBUser"`
	// DBPassword      string `ini:"DBPassword"`
}
type Mail struct {
	Name  string
	Email string
}
type Gps_data struct {
	// Id              bson.ObjectId "_id"
	D_id            string
	I_no            string
	Gps_wd          float64
	Gps_jd          string
	Gps_speed       string
	Gps_real_dir    string
	Gps_dir_extends string
	G_time          string
	C_time          string
	D_version       string
	Gdgps_wd        string
	Gdgps_jd        string
}

// type Gps_data struct {
// 	_id             bson.ObjectId
// 	d_id            string
// 	i_no            string
// 	gps_wd          string
// 	gps_jd          string
// 	gps_speed       string
// 	gps_real_dir    string
// 	gps_dir_extends string
// 	d_time          string
// 	c_time          string
// 	d_version       string
// 	gdgps_wd        string
// 	gdgps_jd        string
// }

func readmongodb() {
	config, err := ReadConfig(filepath) //也可以通过os.arg或flag从命令行指定配置文件路径
	if err != nil {
		log.Fatal(err)
	}
	session, err1 := mgo.Dial(config.MongoDb_hostname)
	if err != nil {
		log.Fatal(err1)
		return
	}
	log.Println("Connect successfully!")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("gps_location") //数据库名称
	collection := db.C("gps_data")   //如果该集合已经存在的话，则直接返回
	if collection == nil {
		log.Println("Please create sheet firstly!")
	}

	// {
	// 	"_id" : ObjectId("59c395e167914b2de8247ee6"),
	// 	"d_id" : "868120174576584",
	// 	"i_no" : "-1",
	// 	"gps_wd" : 29.280464,
	// 	"gps_jd" : 120.038916,
	// 	"gps_speed" : NumberInt(6),
	// 	"gps_real_dir" : NumberInt(317),
	// 	"gps_dir_extends" : NumberInt(53),
	// 	"d_time" : ISODate("2017-09-21T10:35:14.000+0000"),
	// 	"c_time" : ISODate("2017-09-21T10:35:13.749+0000"),
	// 	"d_version" : NumberInt(1),
	// 	"gdgps_wd" : 29.277895507813,
	// 	"gdgps_jd" : 120.04361029731
	// }
	// "_id": bson.NewObjectId(),
	err = collection.Insert(bson.M{"d_id": "868120174576584",
		"i_no":            "-1",
		"gps_wd":          29.280464,
		"gps_jd":          "120.038916",
		"gps_speed":       "6",
		"gps_real_dir":    "317",
		"gps_dir_extends": "53",
		"d_time":          "2017-09-21T10:35:14.000+0000",
		"c_time":          "2017-09-21T10:35:13.749+0000",
		"d_version":       "1",
		"gdgps_wd":        "29.277895507813",
		"gdgps_jd":        "120.04361029731"})
	// m1 := Mail{bson.NewObjectId(), "user1", "user1@dotcoo.com"}
	// err = collection.Insert(&m1)

	if err != nil {
		log.Println("Failed")
	}

	result := Gps_data{}
	err = collection.Find(bson.M{"d_id": "868120174576584"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.D_id, result.Gps_wd, result.C_time)
}

//读取配置文件并转成结构体
func ReadConfig(path string) (*Config, error) {
	//var config Config
	conf, err := ini.Load(path) //加载配置文件
	if err != nil {
		log.Println("load config file fail!")
		os.Exit(1)
	}

	// 映射一个分区
	config := new(Config)
	conf.BlockMode = false
	err = conf.Section("Mongodb").MapTo(config) //解析成结构体
	// log.Println(conf.Section("Mongodb").Key("MongoDb_hostname").String())
	if err != nil {
		log.Println("mapto config file fail!")
		os.Exit(1)
	}
	// log.Println(config.MongoDb_hostname)

	return config, nil
}
