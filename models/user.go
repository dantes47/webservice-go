package models

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
	us.ID = nextID
	nextID++
	users = append(users, &us)
	return us, nil
}
