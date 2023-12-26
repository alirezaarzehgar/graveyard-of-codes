package database

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func Init() (*mgo.Database, error) {
	s, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return nil, fmt.Errorf("faild to dial mongodb: %w", err)
	}
	return s.DB("appdb"), nil
}
