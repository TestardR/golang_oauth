package oauth

import "example.com/utils/errors"

var (
	users = map[string]*User{
		"romain": {Id: 123, Username: "romain"},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("no user found with given parameters")
	}
	return user, nil
}
