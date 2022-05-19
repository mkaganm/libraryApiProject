package db

import (
	"database/sql"
	"fmt"
	"log"
	"model"
	"strings"

	_ "github.com/lib/pq"
)

func InsertAuthor(authorname string) int {
	Initialization()
	defer CloseDB()

	authorname = strings.ToUpper(authorname)

	result, _ := DB.Exec("INSERT INTO authors(authorname) VALUES($1)", authorname)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Effected (%d)", rowsAffected)

	// ! LastInsertId is not supported by this driver
	// lastID, mkm := result.LastInsertId()
	// fmt.Println("\nLAST ID", lastID, mkm)

	lastID, err := GetAuthorIDByName(authorname)
	CheckErr(err)

	return lastID
}

func GetAuthorNameByID(id int) (string, error) {
	Initialization()
	defer CloseDB()

	var authorn string

	err := DB.QueryRow("SELECT authorname FROM authors WHERE authorid=$1", id).Scan(&authorn)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return authorn, err
	case err != nil:
		log.Fatal(err)
		return authorn, err
	default:
		fmt.Printf("AuthorName %s\n", authorn)
		return authorn, nil
	}
}

func GetAuthorIDByName(authorn string) (int, error) {
	Initialization()
	defer CloseDB()

	authorn = strings.ToUpper(authorn)

	var id int

	err := DB.QueryRow("SELECT authorid FROM authors WHERE authorname=$1", authorn).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORDS")
		return id, err
	case err != nil:
		log.Fatal(err)
		return id, err
	default:
		fmt.Printf("ID %d\n", id)
		return id, nil
	}
}

func InsertCategory(categoryname string) int {
	Initialization()
	defer CloseDB()

	categoryname = strings.ToUpper(categoryname)

	result, _ := DB.Exec("INSERT INTO categories(categoryname) VALUES($1)", categoryname)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Effected (%d)", rowsAffected)

	lastID, err := GetCategoryIDByName(categoryname)
	CheckErr(err)

	return lastID
}

func GetCategoryNameByID(id int) (string, error) {
	Initialization()
	defer CloseDB()

	var categoryn string

	err := DB.QueryRow("SELECT categoryid FROM categories WHERE categoryid=$1", id).Scan(&categoryn)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return categoryn, err
	case err != nil:
		log.Fatal(err)
		return categoryn, err
	default:
		fmt.Printf("CategoryName %s\n", categoryn)
		return categoryn, nil
	}
}

func GetCategoryIDByName(categoryn string) (int, error) {
	Initialization()
	defer CloseDB()

	var id int

	categoryn = strings.ToUpper(categoryn)

	err := DB.QueryRow("SELECT categoryid FROM categories WHERE categoryname=$1", categoryn).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return id, err
	case err != nil:
		log.Fatal(err)
		return id, err
	default:
		fmt.Printf("CategoryID %d\n", id)
		return id, nil
	}

}

func InsertPublisher(publishername string) int {
	Initialization()
	defer CloseDB()

	publishername = strings.ToUpper(publishername)

	result, _ := DB.Exec("INSERT INTO publishers(publishername) VALUES($1)", publishername)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Effected (%d)", rowsAffected)

	lastID, err := GetPublisherIDByName(publishername)
	CheckErr(err)

	return lastID
}

func GetPublisherNameByID(id int) (string, error) {
	Initialization()
	defer CloseDB()

	var publishern string

	err := DB.QueryRow("SELECT publishername FROM publishers WHERE publisherid=$1", id).Scan(&publishern)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return publishern, err
	case err != nil:
		log.Fatal(err)
		return publishern, err
	default:
		fmt.Printf("PublisherName %s\n", publishern)
		return publishern, nil
	}
}

func GetPublisherIDByName(publishern string) (int, error) {
	Initialization()
	defer CloseDB()

	var id int
	publishern = strings.ToUpper(publishern)

	err := DB.QueryRow("SELECT publisherid FROM publishers WHERE publishername=$1", publishern).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return id, err
	case err != nil:
		log.Fatal(err)
		return id, err
	default:
		fmt.Printf("PublisherID %d", id)
		return id, nil
	}
}

func InsertBook(book model.Book) error {
	Initialization()
	defer CloseDB()

	book.BookName = strings.ToUpper(book.BookName)

	result, err := DB.Exec("INSERT INTO books(bookname, bookamount, bookauthorid, bookcategoryid, bookpublisherid) VALUES($1, $2, $3, $4, $5)", book.BookName, book.BookAmount, book.BookAuthorID, book.BookCategoryID, book.BookPublisherID)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Effected (%d)", rowsAffected)

	return err
}

func GetBooks() []*model.Book {
	Initialization()
	defer CloseDB()

	rows, err := DB.Query("SELECT *FROM books")

	if err == sql.ErrNoRows {
		fmt.Println("NO RECORDS FOUND!")
	}
	CheckErr(err)

	defer rows.Close()

	var books []*model.Book

	for rows.Next() {
		bk := &model.Book{}
		err := rows.Scan(&bk.BookID, &bk.BookName, &bk.BookAmount, &bk.BookAuthorID, &bk.BookCategoryID, &bk.BookPublisherID)
		CheckErr(err)

		books = append(books, bk)
	}

	err = rows.Err()
	CheckErr(err)

	fmt.Println("\n----------BOOK TABLE--------------")

	for _, bks := range books {
		fmt.Printf("%d\t%s\t%s\t%d\t%d\t%d\n", bks.BookID, bks.BookName, bks.BookAmount, bks.BookAuthorID, bks.BookCategoryID, bks.BookPublisherID)
	}

	return books
}

func GetBookAllByID(id int) model.Book {
	Initialization()
	defer CloseDB()

	var book model.Book

	err := DB.QueryRow("SELECT *FROM books WHERE bookid=$1", id).Scan(&book.BookID, &book.BookName, &book.BookAmount, &book.BookAuthorID, &book.BookCategoryID, &book.BookPublisherID)
	CheckErr(err)

	return book

}

func GetBookIDByName(bookname string) (int, error) {
	Initialization()
	defer CloseDB()

	var id int

	bookname = strings.ToUpper(bookname)

	err := DB.QueryRow("SELECT bookid FROM books WHERE bookname=$1", bookname).Scan(&id)

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

func InsertBooksAndUsers(bookid, userid int) error {
	Initialization()
	defer CloseDB()

	result, err := DB.Exec("INSERT INTO booksandusers(bookid, userid) VALUES($1, $2)", bookid, userid)

	lastinserid, _ := result.LastInsertId()
	fmt.Printf("ID (%d)", lastinserid)

	return err
}

func InsertTag(tag string, bookid int) error {
	Initialization()
	defer CloseDB()

	tag = strings.ToUpper(tag)

	result, err := DB.Exec("INSERT INTO tags(bookid, tags) VALUES($1, $2)", bookid, tag)
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Effected (%d)", rowsAffected)

	return err

}

func GetBookIDByTag(tag string) int {
	Initialization()
	defer CloseDB()

	var id int

	err := DB.QueryRow("SELECT bookid FROM tags WHERE tags=$1", tag).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("NO RECORD")
		return id
	case err != nil:
		log.Fatal(err)
		return id
	default:
		fmt.Printf("ID %d\n", id)
		return id
	}

}
