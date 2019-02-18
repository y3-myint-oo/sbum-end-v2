package utils

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	m "ymo.me/sbum-end-v2/models"
)

// StocksCollection -  Data Model CRUD operation
var StocksCollection *mgo.Collection

// FindAllStock - All Avaliable Unit measure in system
func FindAllStock() ([]m.Stock, error) {
	var datas []m.Stock
	err := StocksCollection.Find(bson.M{}).All(&datas)
	return datas, err
}

// FindAllStockBySerial - All Avaliable Unit measure in system
func FindAllStockBySerial(serial string) (m.Stock, error) {
	var data m.Stock
	//c.Find(bson.M{"name": id}).One(&profile)
	err := StocksCollection.Find(bson.M{"serial": serial}).One(&data)
	return data, err
}

// InsertStock - Create new Unit
func InsertStock(unit m.Stock) error {
	err := StocksCollection.Insert(&unit)
	return err
}

// DeleteStock - Delete Unit
func DeleteStock(data m.Stock) error {
	//cData, err := FindAllStockBySerial(data.Serial)
	//	if err != nil {
	//	return err
	//	}
	//coll.Remove(bson.M{"nm": "Stewart Copeland"})
	err := StocksCollection.Remove(bson.M{"serial": data.Serial})
	return err
}

// UpdateStock -
func UpdateStock(data m.Stock) error {
	//collection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"name": "new Name"}}
	return errors.New("")
}
