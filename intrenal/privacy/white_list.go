package privacy

type UserID int64

type UserData struct {
	Id UserID
    Name string //optional
    // TODO add timepoint ?
}

type WhiteListEntry UserData

type WhiteUserList interface {
    add(user UserID) error
    FindByID(user UserID) (WhiteListEntry, error)
    Remove(user UserID) error
}
