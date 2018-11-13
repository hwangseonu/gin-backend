package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

var DB *mgo.Database

func init() {
	if s, err := mgo.Dial(os.Getenv("MONGODB_URI")); err != nil {
		panic(err)
	} else {
		DB = s.DB(os.Getenv("dbname"))
	}
}

func GetNextId(doc string) int {
	var counter map[string]interface{}

	if err := DB.C("identitycounters").Find(bson.M{"document": doc}).One(&counter); err != nil || counter == nil {
		DB.C("identitycounters").Insert(map[string]interface{}{
			"count": 1,
			"document": doc,
		})
		return 0
	} else {
		id := counter["count"].(int)
		counter["count"] = id + 1
		DB.C("identitycounters").Update(bson.M{"document": doc}, counter)
		return id
	}
}