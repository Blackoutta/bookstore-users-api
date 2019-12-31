package users

import (
	"fmt"
	"github.com/Blackoutta/bookstore-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)


func (u *User) Save() *errors.RestErr {
	current := usersDB[u.ID]
	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists",  u.ID))
	}
	usersDB[u.ID] = u
	return nil
}

func (u *User) Get() *errors.RestErr {
	result := usersDB[u.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.ID))
	}
	u.ID = result.ID
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated
	return nil
}