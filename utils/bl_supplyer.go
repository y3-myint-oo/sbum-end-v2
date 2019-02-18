package utils

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	m "ymo.me/sbum-end-v2/models"
)

// SupplyerCollection -  Data Model CRUD operation
var SupplyerCollection *mgo.Collection

// FindAllSupplyers - All Avaliable Unit measure in system
func FindAllSupplyers() ([]m.Supply, error) {
	var datas []m.Supply
	err := SupplyerCollection.Find(bson.M{}).All(&datas)
	return datas, err
}

// FindAllSupplyerByName - All Avaliable Unit measure in system
func FindAllSupplyerByName(name string) (m.Supply, error) {
	var data m.Supply
	//c.Find(bson.M{"name": id}).One(&profile)
	err := SupplyerCollection.Find(bson.M{"name": name}).One(&data)
	return data, err
}

// FindAllSupplyerByPhone - All Avaliable Unit measure in system
func FindAllSupplyerByPhone(phone string) (m.Supply, error) {
	var data m.Supply
	//c.Find(bson.M{"name": id}).One(&profile)
	err := SupplyerCollection.Find(bson.M{"phone": phone}).One(&data)
	return data, err
}

// InsertSupplyer - Create new Unit
func InsertSupplyer(unit m.Supply) error {
	err := SupplyerCollection.Insert(&unit)
	return err
}

// DeleteSupplyer - Delete Unit
func DeleteSupplyer(data m.Stock) error {
	//cData, err := FindAllStockBySerial(data.Serial)
	//	if err != nil {
	//	return err
	//	}
	//coll.Remove(bson.M{"nm": "Stewart Copeland"})
	err := SupplyerCollection.Remove(bson.M{"serial": data.Serial})
	return err
}

// UpdateSupplyer -
func UpdateSupplyer(data m.Stock) error {
	//collection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"name": "new Name"}}
	return errors.New("")
}
