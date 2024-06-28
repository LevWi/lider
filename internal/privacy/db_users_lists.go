package privacy

func InitWhiteList() (DBUsersList, error) {
	return initUserList(defaultDbPath, "white_list")
}

func InitWaitingList() (DBUsersList, error) {
	return initUserList(defaultDbPath, "waiting_list")
}