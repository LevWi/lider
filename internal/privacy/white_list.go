package privacy

import "errors"

type UserID int64

type UserData struct {
	Id UserID
    Name string
    // TODO add timepoint ?
}

type WhiteListEntry UserData

type WhiteUserList interface {
    Add(user WhiteListEntry) error
    FindByID(userId UserID) (WhiteListEntry, error)
    Remove(userId UserID) error
}

var ErrNotFound = errors.New("WhiteUserList: not found")