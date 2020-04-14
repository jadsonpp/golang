package models

// User Struct
type User struct {
	ID        int
	Firstname string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// Get all users
func GetUsers() []*User {
	return users
}

//Add a new User
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
