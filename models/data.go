package models

import "gopkg.in/mgo.v2/bson"

// Data - Sample
type Data struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `json:"title"`
}

// MockDatas - Mock data set for Todo
var MockDatas = []Data{
	Data{"507f1f77bcf86cd799439011", "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Data{"507f1f77bcf86cd799439012", "Why did the coffee file a police report? It got mugged."},
	Data{"507f1f77bcf86cd799439013", "How does a penguin build it's house? Igloos it together."},
}
