package privacy

type UserID int64

type UserData struct {
	Id UserID
    Name string //optional
    // TODO add timepoint ?
}

type WhiteListEntry UserData

type WhiteUserList interface {
    Add(user WhiteListEntry) error
    FindByID(userId UserID) (WhiteListEntry, error)
    Remove(userId UserID) error
}
