package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rithium/stor-auth/config"
	"log"
	"github.com/go-sql-driver/mysql"
)

type Datastore interface {
	CreateApiKey()(*ApiKey, error)
	KeyExists(key string)(bool, error)
	FindActiveKey(key string)(*ApiKey, error)
}

type DB struct {
	*sql.DB
}

func NewDb(params config.MySQLConfig)(*DB, error) {
	conf := mysql.Config{User: params.User, Passwd: params.Pass, DBName: params.Database, Net: "tcp", Addr: params.Url+":"+params.Port}

	log.Printf("%+v", conf)

	db, err := sql.Open("mysql", conf.FormatDSN())

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}