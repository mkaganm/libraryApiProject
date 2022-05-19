package db

import (
	"database/sql"
	"fmt"
	"log"
	"model"

	_ "github.com/lib/pq"
)

func InsertUser(user model.User) error {
	Initialization()
	defer CloseDB()

	result, err := DB.Exec("INSERT INTO users(username, userpassword, firstname, lastname, usermail, userphone) VALUES($1, $2, $3, $4, $5, $6)", user.UserName, user.Password, user.FirstName, user.LastName, user.Email, user.Phone)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Effected (%d)", rowsAffected)

	// ! LastInsertId is not supported by this driver
	// lastID, mkm := result.LastInsertId()
	// fmt.Println("\nLAST ID", lastID, mkm)

	return err
}

func GetUser() {
	Initialization()
	defer CloseDB()

	rows, err := DB.Query("SELECT *FROM users")

	if err == sql.ErrNoRows {
		fmt.Println("NO RECORDS FOUND!")
	}
	CheckErr(err)

	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		usr := &model.User{}
		err := rows.Scan(&usr.ID, &usr.UserName, &usr.Password, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Phone)
		CheckErr(err)
		users = append(users, usr)
	}

	err = rows.Err()
	CheckErr(err)
	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("\n--------USER TABLE-------")
	for _, us := range users {
		fmt.Printf("%d\t%s\t%s\t%s\t%s\t%s\t%s\t\n", us.ID, us.UserName, us.Password, us.FirstName, us.LastName, us.Email, us.Phone)

	}

}

func GetUserAllByID(id int) []*model.User {
	Initialization()
	defer CloseDB()

	rows, err := DB.Query("SELECT *FROM users WHERE userid=$1", id)

	if err == sql.ErrNoRows {
		fmt.Println("NO RECORDS FOUND!")
	}
	CheckErr(err)

	defer rows.Close()

	var user []*model.User
	for rows.Next() {
		usr := &model.User{}
		err := rows.Scan(&usr.ID, &usr.UserName, &usr.Password, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Phone)
		CheckErr(err)
		user = append(user, usr)
	}

	err = rows.Err()
	CheckErr(err)

	return user
}

func GetUsernameByID(id int) (string, error) {
	Initialization()
	defer CloseDB()

	var usern string

	err := DB.QueryRow("SELECT username FROM users WHERE userid=$1", id).Scan(&usern)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return usern, err
	case err != nil:
		log.Fatal(err)
		return usern, err
	default:
		fmt.Printf("Username %s\n", usern)
		return usern, nil
	}
}

func GetPassByID(id int) (string, error) {
	Initialization()
	defer CloseDB()

	var userp string

	err := DB.QueryRow("SELECT userpassword FROM users WHERE userid=$1", id).Scan(&userp)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return userp, err
	case err != nil:
		log.Fatal(err)
		return userp, err
	default:
		fmt.Printf("Password %s\n", userp)
		return userp, nil
	}
}

func GetIDByUsername(name string) (int, error) {
	Initialization()
	defer CloseDB()

	var userID int

	err := DB.QueryRow("SELECT userid FROM users WHERE username=$1", name).Scan(&userID)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return userID, err
	case err != nil:
		log.Fatal(err)
		return userID, err
	default:
		fmt.Printf("ID %d\n", userID)
		return userID, nil
	}
}
