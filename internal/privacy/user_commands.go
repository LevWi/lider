package privacy

import "errors"

type UserCommands struct {
	whiteList   UsersList
	waitingList UsersList
}

var ErrInWaitingList = errors.New("user in waiting list")
var ErrAddToWaitingList = errors.New("user add to waiting list")

func (uc *UserCommands) GrantedAccessCheck(id UserID) (UserData, error) {
	ud, err := uc.whiteList.FindByID(id)
	if err == ErrNotFound {
		ud, err = uc.waitingList.FindByID(id)
		if err == nil {
			err = ErrInWaitingList
		} else if err == ErrNotFound {
			err = uc.waitingList.Add(UsersListEntry{Id: id})
			if err == nil {
				err = ErrAddToWaitingList
			}
			ud.Id = id
		}
	}

	return UserData(ud), err
}
