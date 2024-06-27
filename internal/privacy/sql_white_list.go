package privacy

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName        = "white_list"
	defaultDbPath = dbName + ".db"
	dbDriver      = "sqlite3"
	tbKeyId       = "id"
	tbKeyName     = "name"

	sqlInsertPattern = "INSERT INTO " + dbName + "(" + tbKeyId + ", " + tbKeyName + ") VALUES (?, ?)"
	sqlFindPattern   = "SELECT * FROM " + dbName + " WHERE " + tbKeyId + " = ?"
	sqlDeletePattern = "DELETE FROM " + dbName + " WHERE " + tbKeyId + " = ?"
)

type WhiteUserListDB struct {
	*sql.DB
	//TODO add cache
}

func InitWhiteList() (WhiteUserListDB, error) {
	return initWhiteList(defaultDbPath)
}

func initWhiteList(dbName string) (WhiteUserListDB, error) {
	db, err := sql.Open(dbDriver, dbName)
	if err != nil {
		return WhiteUserListDB{}, err
	}

	_, err = db.Exec(fmt.Sprintf(`create table if not exists white_list (
        %s  integer not null,
        %s  text unique,
        primary key(%s)
    );`, tbKeyId, tbKeyName, tbKeyId))
	if err != nil {
		close_res := db.Close()
		return WhiteUserListDB{}, errors.Join(err, close_res)
	}
	return WhiteUserListDB{db}, nil
}

func (db WhiteUserListDB) Add(data WhiteListEntry) error {
	_, err := db.Exec(sqlInsertPattern, data.Id, data.Name)
	return err
}

func (db WhiteUserListDB) FindByID(userId UserID) (out WhiteListEntry, err error) {
	err = db.QueryRow(sqlFindPattern, userId).Scan(&out.Id, &out.Name)
	if err == sql.ErrNoRows {
		err = ErrNotFound
	}
	return
}

func (db WhiteUserListDB) Remove(userId UserID) error {
	_, err := db.Exec(sqlDeletePattern, userId)
	return err
}
