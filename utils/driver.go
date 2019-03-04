package utils

import (
	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

const (
	DTATBASE = "test003"
)

var DB *mgo.Database

// InitDriver - Init the MongoDb Connection
func InitDriver() error {
	//mongodb://mongo:27018
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	DB = session.DB(DTATBASE)
	DatasCollection = DB.C("data")
	UnitCollection = DB.C("unit")
	StocksCollection = DB.C("stock")
	SupplyerCollection = DB.C("supplyer")
	return nil
}
