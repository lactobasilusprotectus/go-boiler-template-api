package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lactobasilusprotectus/go-template/config"
	"log"
)

type DatabaseConnection struct {
	Db *sqlx.DB
}

func New(config config.DatabaseConfig) (*DatabaseConnection, error) {
	env := config.Driver
	log.Println(env)
	var db *sqlx.DB

	switch env {
	case "mysql":
		mysql, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Database))

		if err != nil {
			log.Println("DB Connection err: ", err)
			return nil, err
		}
		mysql.SetMaxIdleConns(config.MaxIdleConnection)
		mysql.SetMaxOpenConns(config.MaxOpenConnection)

	}

	return &DatabaseConnection{
		Db: db,
	}, nil
}

func closeConnection(db *DatabaseConnection) {
	_ = db.Db.Close()
}
