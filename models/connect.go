package models

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"errors"
)

type Map bson.M

var sessions map[string]*mgo.Session

// GetSession - get a db session handle
func GetSession(db string) *mgo.Session {
	initSessions()
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"ds157639.mlab.com:57639"},
	 	Timeout:  60 * time.Second,
		Database: db,
		Username: "bui",
		Password: "bui123",
	}
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetSafe(&mgo.Safe{WMode: "majority"})
	return mongoSession
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

func Find(db string, collection string, query Map) []Map {
	result := make([]Map, 0)
	session := GetSession(db)
	defer session.Close()
	coll := session.DB(db).C(collection)
	coll.Find(query).All(&result)
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

func Delete(db string, collection string, query Map) error {
	session := GetSession(db)
	defer session.Close()
	res := Find(db,collection,query);
	coll := session.DB(db).C(collection)
	for _, b := range res {
	err := coll.Remove(b)
		if err != nil {
			log.Println(err)
		}
	}
	err1 := errors.New("")
	return err1
}
