package db

import (
	"database/sql"
	"fmt"
	"log"
	"model"

	_ "github.com/lib/pq"
)

func InsertAdmin(admin model.Admin) {
	Initialization()
	defer CloseDB()

	result, _ := DB.Exec("INSERT INTO admins(adminname, adminpassword) VALUES($1, $2)", admin.AdminName, admin.Password)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Effected (%d)", rowsAffected)
}

func GetAdmin() {
	Initialization()
	defer CloseDB()

	rows, err := DB.Query("SELECT *FROM admins")
	CheckErr(err)
	if err == sql.ErrNoRows {
		fmt.Println("NO RECORDS")
	}

	defer rows.Close()

	var admins []*model.Admin

	for rows.Next() {
		adm := &model.Admin{}
		err := rows.Scan(&adm.ID, &adm.AdminName, &adm.Password)
		CheckErr(err)
		admins = append(admins, adm)
	}

	err = rows.Err()
	CheckErr(err)
	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("\n-------ADMIN TABLE--------")
	for _, ad := range admins {
		fmt.Printf("%d\t%s\t%s\n", ad.ID, ad.AdminName, ad.Password)
	}
}

func GetAdminNameByID(id int) (string, error) {
	Initialization()
	defer CloseDB()

	var adminn string

	err := DB.QueryRow("SELECT adminname FROM admins WHERE adminid=$1", id).Scan(&adminn)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return adminn, err
	case err != nil:
		log.Fatal(err)
		return adminn, err
	default:
		fmt.Printf("AdminName %s\n", adminn)
		return adminn, nil
	}
}

func GetIDByAdminname(adminn string) (int, error) {
	Initialization()
	defer CloseDB()

	var id int

	err := DB.QueryRow("SELECT adminid FROM admins WHERE adminname=$1", adminn).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return id, err
	case err != nil:
		log.Fatal(err)
		return id, err
	default:
		fmt.Printf("ID %d\n", id)
		return id, nil
	}
}

func GetAdminPassByID(id int) (string, error) {
	Initialization()
	defer CloseDB()

	var password string

	err := DB.QueryRow("SELECT adminpassword FROM admins WHERE adminid=$1", id).Scan(&password)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return password, err
	case err != nil:
		log.Fatal(err)
		return password, err
	default:
		fmt.Printf("Password %s\n", password)
		return password, nil
	}
}
