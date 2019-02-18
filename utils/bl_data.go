package utils

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	m "ymo.me/sbum-end-v2/models"
)

// DatasCollection -  Data Model CRUD operation
var DatasCollection *mgo.Collection

// FindAllData - All Avaliable Unit measure in system
func FindAllData() ([]m.Data, error) {
	var datas []m.Data
	err := DatasCollection.Find(bson.M{}).All(&datas)
	return datas, err
}

// InsertData - Create new Unit
func InsertData(unit m.Data) error {
	err := DatasCollection.Insert(&unit)
	return err
}

// DeleteData - Delete Unit
func DeleteData(data m.Data) error {
	//Make Sure at least one unit have
	err := DatasCollection.Remove(&data)
	return err
}

// UpdateData -
func UpdateData(data m.Data) error {
	//collection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"name": "new Name"}}
	return errors.New("")
}
