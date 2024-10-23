package user

import "errors"

type Person struct {
	Username string
	Password string
}

type Users struct {
	Users []Person
}

func Login(uname, pass string, users Users) error {
	for _, user := range users.Users {
		if uname == "" || pass == "" {
			return errors.New("Username and Password is required")
		}
		if user.Username == uname && user.Password == pass {
			break
		} else {
			return errors.New("Invalid username or password")
		}
	}
	return nil
}

func AddUsers(p Person, u Users) Users {
	u.Users = append(u.Users, p)
	return u
}
