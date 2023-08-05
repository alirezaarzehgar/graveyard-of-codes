package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConf struct {
	Username string
	Password string
	Hostname string
	Port     string
	DbName   string
}

func NewConnection(conf MysqlConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Username, conf.Password, conf.Hostname, conf.Port, conf.DbName)
	log.Printf("conf:%#v\n", conf)
	log.Println("dsn:", dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	sqlDb, _ := db.DB()
	err = sqlDb.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return db, nil
}
