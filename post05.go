package post05

import (
	_ "github.com/lib/pq"
)

// Connection details

type User struct {
	ID       int
	Username string
}

type Userdata struct {
	ID          int
	Name        string
	Surname     string
	Description string
}

func initiateConnection() {

}

// The function returns the User ID of the user
// -1 if the user does not exist
func exists(id int) int {

}

// AddUser adds a new user to the database
func AddUser() bool {

	return true
}

// DeleteUser deletes an existing user
func DeleteUser(id int) bool {

	return true
}

func ListUsers() {

}

func UpdateUser() bool {

	return true
}
