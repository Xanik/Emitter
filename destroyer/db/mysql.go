package db

import (
	"log"

	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB is database session container
type DB struct {
	c *sqlx.DB
}

// Permanent session
var cdb *DB

var dbLock sync.Mutex

// Connect will initialise a permanent SQL session
func Connect() (*DB, error) {
	if cdb == nil {
		fmt.Printf("Creating new DB connection")
		newDsn := fmt.Sprintf("host=postgres_db port=5432 user=root dbname=event password=mysql21 sslmode=disable")
		conn, err := sqlx.Connect("postgres", newDsn)
		if err != nil {
			log.Printf("Could not connect to db: %v, %v", err, newDsn)
			return nil, err
		}
		cdb = &DB{c: conn}
	}
	err := cdb.Reconnect()
	if err != nil {
		fmt.Errorf("Could not reconnect to db")
		return nil, err
	}
	return cdb, nil
}

// Reconnect will try to reconnect to database if connection is lost
func (db *DB) Reconnect() error {
	err := db.c.Ping()
	if err != nil {
		fmt.Printf("Reconnecting to db")
		newDsn := fmt.Sprintf("host=postgres_db port=5432 user=root dbname=event password=mysql21 sslmode=disable")
		c, err := sqlx.Connect("mysql", newDsn)
		if err != nil {
			db.c = nil
			fmt.Errorf("Could not reconnect to db")
			return err
		}
		db.c = c
	}
	return nil
}

type Target struct {
	ID        string `json:"id" db:"id"`
	Message   string `json:"message" db:"message"`
	CreatedOn string `json:"created_on" db:"created_on"`
}

// GetAllTargets will return list of targets from database
func (db *DB) GetAllTargets(targetl string) ([]Target, error) {
	err := db.Reconnect()
	if err != nil {
		return nil, err
	}
	var targets []Target
	if targetl != "*" {
		err = db.c.Select(&targets, "select id, `message`, `created_on` from targets where id=?", targetl)
		if err != nil {
			fmt.Errorf("Can't list targets %v", err)
		}
	} else {
		err = db.c.Select(&targets, "select id, `message`, `created_on` from targets")
		if err != nil {
			fmt.Errorf("Can't list targets %v", err)
		}
	}
	return targets, err
}
