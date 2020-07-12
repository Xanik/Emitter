package test

import (
	"errors"
	"fmt"
	"testing"
)

type fakeDB struct {
	T *testing.T
}

type TestDBListTarget struct {
	ID        string
	Message   string
	CreatedOn string
}

func (db fakeDB) Select(dest interface{}, query string, args ...interface{}) error {
	if query == "select id, `message`, created_on from targets where id=?" {
		loc := dest.(*[]TestDBListTarget)
		switch args[0] {
		case "1":
			*loc = []TestDBListTarget{
				{
					ID:        "01EBP4DP4VECW8PHDJJFNEDVKE",
					Message:   "some message to send",
					CreatedOn: "2020-06-25T16:23:37.720Z",
				},
			}
		case "2":
			*loc = []TestDBListTarget{
				{
					ID:        "20EBP4DQBVECW8PHDJJFNEDVKE",
					Message:   "some other message to send",
					CreatedOn: "2020-06-25T16:23:37.720Z",
				},
			}
		}
	} else {
		db.T.Errorf("Invalid query parameter %s", query)
	}
	return nil
}

func (db fakeDB) ListTarget(targetl string) ([]TestDBListTarget, error) {
	var target []TestDBListTarget
	var err error

	err = db.Select(&target, "select id, message, created_on from targets where id=$1", targetl)

	if err != nil {
		e := fmt.Sprintf("No such Target: with ID '%s' %v", targetl, err)
		return nil, errors.New(e)
	}

	return target, nil
}
