package models

import "gopkg.in/mgo.v2/bson"

// Stock - Warehouse Item
type Stock struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Serial     string        `json:"serial"`
	Name       string        `json:"title"`
	Unit       string        `json:"unit"`
	Price      []StockPrice  `json:"price"`
	Network    string        `json:"network"` // pending,synced
	TotalPrice string        `json:"totalPrice"`
}

// StockPrice - WareHouse Item price history
type StockPrice struct {
	Date   string `json:"date"`
	SPrice string `json:"sprice"`
	BPrice string `json:"bprice"`
}

// MockStocks -
var MockStocks = []Stock{}
