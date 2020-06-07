package models

import (
	"errors"
	"fmt"
)

// User is ..
type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

var (
	users  []*User
	nextID = 1
)

// GetUsers is..
func GetUsers() []*User {
	return users
}

// AddUsers is..
func AddUsers(us User) (User, error) {
	if us.ID != 0 {
		return User{}, errors.New("New Users must not include ID or it must be set to 0")
	}
	us.ID = nextID
	nextID++
	users = append(users, &us)
	return us, nil
}

// GetUserByID is..
func GetUserByID(id int) (User, error) {
	for _, us := range users {
		if us.ID == id {
			return *us, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser is..
func UpdateUser(us User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == us.ID {
			users[i] = &us
			return us, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", us.ID)
}

// RemoveUserByID is..
func RemoveUserByID(id int) error {
	for i, us := range users {
		if us.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
