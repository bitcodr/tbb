package config

import "github.com/globalsign/mgo"

//DBDial connect to mongodb
func DBDial() (*mgo.Session, error) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		return nil, err
	}
	return session, nil
}
