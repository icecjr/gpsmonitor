package main

import (
	"fmt"
	. "project/elangshen/mutils"

	mgo "gopkg.in/mgo.v2"
)

//var DbSession *mgo.Session
//var GpsDatabase *mgo.Database

type DB struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

func (db *DB) ConnectMongoDb() {
	//prepare mongodb *************************************************
	var err error
	db.Session, err = mgo.Dial(mongodbConfig.MongoDb_hostname)
	CheckError(err)

	// Optional. Switch the session to a monotonic behavior.
	db.Session.SetMode(mgo.Monotonic, true)
	fmt.Println("begin to connect mongodb")
	//GpsDatabase = db.Session.DB(mongodbConfig.DBName)
	//structobj := bson.M{"name": "test1", "Age": "34"}
	//GpsDatabase.C("gps_test").Insert(structobj)
	//db.Insert("gps_test", bson.M{"name": "test1", "Age": "34"})
	//fmt.Println("End of connect mongodb")
	// db.C("test").Insert(bson.M{"name": "3", "value": "5"})
	// log.Println(db.Name)

}
func (db *DB) Insert(sheetName string, structObject interface{}) bool {
	db.SetCollection(sheetName)
	db.Collection.Insert(structObject)
	//CheckError(err)
	return true
}
func (db *DB) CloseMongoDb() {
	db.Session.Close()
}
func (db *DB) SetCollection(collection string) {
	db.Collection = db.Session.DB(mongodbConfig.DBName).C(collection)
}
func (db *DB) Update(collection string, selector interface{}, updator interface{}) {
	db.SetCollection(collection)
	db.Collection.Update(selector, updator)
}
func (db *DB) Find(collection string, selector interface{}, updator interface{}) {
	db.SetCollection(collection)
	db.Collection.Find(selector).One(&updator)
}
