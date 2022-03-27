package db

import (
	"github.com/lactobasilusprotectus/go-template/config"
	"log"
	"testing"
)

func TestNewDb(t *testing.T) {
	connection, err := New(config.DatabaseConfig{})

	log.Println(connection, err)
}
