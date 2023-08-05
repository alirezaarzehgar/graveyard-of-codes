package storage

import (
	"database/sql"
	"fmt"

	"github.com/BaseMax/JokeGoServiceAPI/internal/book"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlStorage struct {
	db *sql.DB
}

type MysqlConfig struct {
	Username string
	Password string
	DbName   string
	Port     uint
	Host     string
}

func NewMysqlStorage(conf MysqlConfig) (MysqlStorage, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return MysqlStorage{}, fmt.Errorf("impossible to open SQL connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return MysqlStorage{}, fmt.Errorf("impossible to pind database: %w", err)
	}
	return MysqlStorage{db}, nil
}

func (s MysqlStorage) Create(b book.Book) error {
	return nil
}
