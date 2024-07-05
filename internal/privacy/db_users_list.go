package privacy

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultDbPath    = "users.db"
	dbDriver         = "sqlite3"
	tbKeyId          = "id"
	tbKeyName        = "name"
)

type DBUsersList struct {
	*sql.DB
	pattrens QueriesPatterns
}

type QueriesPatterns struct {
	Insert string
	Find   string
	Delete string
}

func userListSqlPatterns(tableName string) QueriesPatterns {
	return QueriesPatterns{
		Insert : "INSERT INTO " + tableName + "(" + tbKeyId + ", " + tbKeyName + ") VALUES (?, ?)",
		Find   : "SELECT * FROM " + tableName + " WHERE " + tbKeyId + " = ?",
		Delete : "DELETE FROM " + tableName + " WHERE " + tbKeyId + " = ?",
	}
}

func initUserList(dbName string, tableName string) (DBUsersList, error) {
	db, err := sql.Open(dbDriver, dbName)
	if err != nil {
		return DBUsersList{}, err
	}

	_, err = db.Exec(fmt.Sprintf(`create table if not exists %s (
        %s  integer not null,
        %s  text not null,
        primary key(%s)
    );`, tableName, tbKeyId, tbKeyName, tbKeyId))
	if err != nil {
		close_res := db.Close()
		return DBUsersList{}, errors.Join(err, close_res)
	}
	return DBUsersList{db, userListSqlPatterns(tableName)}, nil
}

func (db *DBUsersList) Add(data UsersListEntry) error {
	_, err := db.Exec(db.pattrens.Insert, data.Id, data.Name)
	//TODO check that already exist
	return err
}

func (db *DBUsersList) FindByID(userId UserID) (out UsersListEntry, err error) {
	err = db.QueryRow(db.pattrens.Find, userId).Scan(&out.Id, &out.Name)
	if err == sql.ErrNoRows {
		err = ErrNotFound
	}
	return
}

func (db *DBUsersList) Remove(userId UserID) error {
	_, err := db.Exec(db.pattrens.Delete, userId)
	return err
}
