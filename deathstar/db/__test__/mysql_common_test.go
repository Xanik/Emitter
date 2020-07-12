package test

import (
	"errors"
	"fmt"
	"testing"
)

type fakeDB struct {
	T *testing.T
}

type TestDBEvent struct {
	ID        string
	Message   string
	CreatedOn string
}

func (db fakeDB) Exec(dest interface{}, query string, args ...interface{}) (int64, error) {
	if query == "" {
		db.T.Error("Invalid query parameter")
	}
	return 1, nil
}

func (db fakeDB) StoreEvents(e TestDBEvent) (int64, error) {
	var err error

	resp, err := db.Exec("INSERT INTO targets (id, message, created_on) VALUES ($1, $2, $3) ", e.ID, e.Message, e.CreatedOn)

	if err != nil {
		e := fmt.Sprintf("Can't insert events %v", err)
		return 0, errors.New(e)
	}

	return resp, nil
}
