package privacy

import (
	//"os"
	"testing"
)

func TestWhiteList(t *testing.T) {
	//dbPath  := t.TempDir() + string(os.PathSeparator)
	dbPath := "test_.db"
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
}
