package models

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Map bson.M

var sessions map[string]*mgo.Session

// GetSession - get a db session handle
func GetSession(db string) *mgo.Session {
	initSessions()
	uri := getConnectionURI(db)
	if sessions[db] == nil {
		log.Println("Initializing Database")
		mgoSession, err := mgo.Dial(uri)
		if err != nil {
			panic(err) // no, not really
		}
		mgoSession.SetSafe(&mgo.Safe{WMode: "majority"})
		sessions[db] = mgoSession
	}
	return sessions[db].Clone()
}

func getConnectionURI(db string) string {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s", "admin", "Welcome1", "35.170.186.108", db)
	return uri
}

func initSessions() {
	if sessions == nil {
		sessions = make(map[string]*mgo.Session)
	}
}

func FindOne(db string, collection string, query Map) Map {
	result := make(Map, 0)
	session := GetSession(db)
	defer session.Close()
	coll := session.DB(db).C(collection)
	coll.Find(query).One(&result)
	return result
}

func Create(database string, collection string, data Map) (Map, error) {
	session := GetSession(database)
	defer session.Close()
	dbCollection := session.DB(database).C(collection)
	err := dbCollection.Insert(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Update(db string, collection string, query Map, upd Map) error {
	session := GetSession(db)
	defer session.Close()
	coll := session.DB(db).C(collection)
	err := coll.Update(query, upd)
	if err != nil {
		log.Println(err)
	}
	return err
}
