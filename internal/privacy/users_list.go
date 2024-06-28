package privacy

import "errors"

type UserID int64

type UserData struct {
	Id UserID
    Name string
}

type UsersListEntry UserData

type UsersList interface {
    Add(user UsersListEntry) error
    FindByID(userId UserID) (UsersListEntry, error)
    Remove(userId UserID) error
}

var ErrNotFound = errors.New("UsersList: not found")