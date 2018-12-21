package config

import (
	"os"

	"github.com/globalsign/mgo"
)

//DBDial connect to mongodb
func DBDial() (*mgo.Session, error) {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		return nil, err
	}
	return session, nil
}
