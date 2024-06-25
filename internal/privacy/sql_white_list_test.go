package privacy

import (
    "database/sql"
	"os"
	"testing"
)

func TestWhiteList(t *testing.T) {
	dbPath := t.TempDir() + string(os.PathSeparator) + "test_.db"
	t.Log("db path:", dbPath)
	db, err := initWhiteList(dbPath)
	if err != nil {
		t.Fatal("initWhiteList fail: ", err)
	}
	defer db.Close()
	dbI := WhiteUserList(db)
	arr := [3]WhiteListEntry{{1234, "Name1"}, {345, "Name2"}, {678, "Name3"}}
	for _, el := range arr {
		err = dbI.Add(el)
		if err != nil {
			t.Fatal(".Add() fail: ", err, el)
		}
	}

	for _, el := range arr {
		user, err := dbI.FindByID(el.Id)
		if err != nil {
			t.Fatal(".FindByID() fail: ", err)
		}
		if user != el {
			t.Fatal(".FindByID() type mismatch: ", user, el)
		}
	}

    checkDeleted := func(userId UserID) {
        user, err := dbI.FindByID(userId)
        if err == nil {
            t.Fatal(".FindByID(). enexpected: ", user)
        }
        if err != sql.ErrNoRows {
            t.Fatal(".FindByID():", err)
        }
    }
    checkDeleted(5)

    err = dbI.Remove(5)
    if err != nil {
		t.Fatal(".Remove(): ", err)
	}

	for _, el := range arr {
		err := dbI.Remove(el.Id)
		if err != nil {
			t.Fatal(".FindByID() fail: ", err)
		}
        checkDeleted(el.Id)
	}
}
