package post05

import (
	"database/sql"
	"fmt"
	"strings"

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
		Hostname, Port, Username, Password, Database)

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
	username = strings.ToLower(username)

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userID := -1
	statement := fmt.Sprintf(`SELECT "id" FROM "users" where username = '%s'`, username)
	fmt.Println("exists()", statement)
	rows, err := db.Query(statement)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("Scan", err)

		}
		userID = id
		fmt.Println("*", userID)
	}
	defer rows.Close()
	return userID
}

// AddUser adds a new user to the database
// Returns new User ID
// -1 if there was an error
func AddUser(d Userdata) int {
	d.Username = strings.ToLower(d.Username)

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userID := exists(d.Username)
	if userID != -1 {
		fmt.Println("User already exists:", Username)
		return -1
	}

	insertStatemet := `insert into "users" ("username") values ($1)`
	_, err = db.Exec(insertStatemet, d.Username)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	userID = exists(d.Username)
	if userID == -1 {
		return userID
	}

	insertStatemet = `insert into "userdata" ("userid", "name", "surname", "description") values ($1, $2, $3, $4)`
	_, err = db.Exec(insertStatemet, userID, d.Name, d.Surname, d.Description)
	if err != nil {
		fmt.Println("db.Exec()", err)
		return -1
	}

	return userID
}

// DeleteUser deletes an existing user
func DeleteUser(id int) bool {
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()

	// Delete from Userdata
	deleteStatement := `delete from "userdata" where id=$1`
	_, err = db.Exec(deleteStatement, id)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Delete from Users
	deleteStatement = `delete from "users" where id=$1`
	_, err = db.Exec(deleteStatement, id)
	if err != nil {
		fmt.Println(err)
		return false
	}

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

	rows, err := db.Query(`SELECT "id","username" FROM "users"`)
	if err != nil {
		fmt.Println("Query", err)
		return Data, err
	}

	for rows.Next() {
		var id string
		var username string
		err = rows.Scan(&id, &username)
		if err != nil {
			fmt.Println("Scan", err)
			return Data, err
		}
		fmt.Println("*", id, username)
	}
	defer rows.Close()

	return Data, nil
}

// UpdateUser is for updating an existing user
// Returns true on success
// False on failure
func UpdateUser(d Userdata) bool {
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()

	return true
}
