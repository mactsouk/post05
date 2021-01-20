package post05

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connection details
var (
	Hostname = ""
	Port     = 2345
	Username = ""
	Password = ""
	Database = ""
)

// User is for holding data from the User table
type User struct {
	ID       int
	Username string
}

// Userdata is for holding full user data
// Userdata table + Username
type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

func openConnection() (*sql.DB, error) {
	// connection string
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Hostname, Post, Username, Password, Database)

	// open database
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// The function returns the User ID of the username
// -1 if the user does not exist
func exists(username string) int {
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	return -1

}

// AddUser adds a new user to the database
func AddUser() bool {
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()

	return true
}

// DeleteUser deletes an existing user
func DeleteUser(id int) bool {
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()

	return true
}

// ListUsers lists all users in the database
func ListUsers() ([]Userdata, error) {
	Data := []Userdata{}

	db, err := openConnection()
	if err != nil {
		return Data, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT "ID","Username" FROM "Users"`)
	if err != nil {
		fmt.Println("Query", err)
		return Data, err
	}

	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("Scan", err)
			return Data, err
		}
		fmt.Println("*", id, name)
	}
	defer rows.Close()

	return Data, nil
}

// UpdateUser is for updating an existing user
// Returns true on success
// False on failure
func UpdateUser() bool {
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()

	return true
}
