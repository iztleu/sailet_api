package mongodb

import (
	"log"

	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
)

func InitDatabaseConnection(host string) error {
	var err error
	session, err = mgo.Dial(host)
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	return nil
}

func GetSession() *mgo.Session {
	if session == nil {
		log.Fatalf("Connection object is nil! This should not happen!")
	}
	return session

}

func CloseDatabaseConnection() {
	if session == nil {
		return
	}
	session.Close()
}
