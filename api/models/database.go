package models

import (
	"os"

	"labix.org/v2/mgo"
)

// Database deals with mongoDB to get connect
func Database() *mgo.Session {
	session, err := mgo.Dial("apidocker_database_1")

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	session.DB(os.Getenv("DB_NAME"))
	return session
}
