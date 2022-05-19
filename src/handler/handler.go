package handler

import (
	"database/sql"
	"db"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"model"
	"net/http"
	"strings"
)

var Tmpl *template.Template

func NotPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func init() {
	Tmpl = template.Must(template.ParseGlob("templates/*"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "index.html", nil)
}

func UserLoginPage(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "userLogin.html", nil)
}

func UserRegisterPage(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "userRegister.html", nil)
}

func AdminLoginPage(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "adminLogin.html", nil)
}

func UserAuth(w http.ResponseWriter, r *http.Request) {

	NotPost(w, r)
	w.Header().Set("Content-Type", "application/json")

	userName := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("Username:" + userName + "\tPassword:" + password)

	usIDDB, err1 := db.GetIDByUsername(userName)
	usPassDB, err2 := db.GetPassByID(usIDDB)

	if err1 == sql.ErrNoRows || err2 == sql.ErrNoRows {
		log.Printf("--------NO RECORD!!!-------")
		str := "SOMETHINGS GOING WRONG"
		jsonstr, _ := json.Marshal(str)
		fmt.Fprintf(w, string(jsonstr))

	} else if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
		str := "SOMETHINGS GOING WRONG"
		jsonstr, _ := json.Marshal(str)
		fmt.Fprintf(w, string(jsonstr))

	} else {
		if usPassDB == password {
			var user []*model.User
			user = db.GetUserAllByID(usIDDB)

			jsonuser, _ := json.Marshal(user)
			fmt.Fprintf(w, string(jsonuser))
		} else {
			str := "SOMETHINGS GOING WRONG"
			jsonstr, _ := json.Marshal(str)
			fmt.Fprintf(w, string(jsonstr))
		}
	}
}

func AdminAuth(w http.ResponseWriter, r *http.Request) {

	NotPost(w, r)

	userName := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("Username:" + userName + "\nPassword:" + password)

	adIDDB, err1 := db.GetIDByAdminname(userName)
	adPassDB, err2 := db.GetAdminPassByID(adIDDB)

	if err1 == sql.ErrNoRows || err2 == sql.ErrNoRows {
		log.Printf("--------NO RECORD!!!-------")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		if adPassDB == password {
			Tmpl.ExecuteTemplate(w, "adminAuth.html", userName)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func RegisterAuth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	NotPost(w, r)

	user := model.User{
		UserName:  r.FormValue("username"),
		Password:  r.FormValue("password"),
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
	}
	fmt.Println(user)
	err := db.InsertUser(user)

	if err != nil {
		log.Fatal(err)
		str := "SOMETHINGS GOING WRONG"
		jsonstr, _ := json.Marshal(str)
		fmt.Fprintf(w, string(jsonstr))
	} else {
		jsonusr, _ := json.Marshal(user)
		fmt.Fprintf(w, string(jsonusr))
	}

}

func BookList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	booklist := db.GetBooks()

	jsonbooklist, _ := json.Marshal(booklist)

	fmt.Fprintf(w, string(jsonbooklist))

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "addBook.html", nil)
}

func BookAuth(w http.ResponseWriter, r *http.Request) {
	NotPost(w, r)
	var err error

	book := model.Book{
		BookName: r.FormValue("bookname"),
		// ! convert to int
		BookAmount: r.FormValue("bookamount"),
	}

	author := model.Author{
		AuthorName: r.FormValue("bookauthor"),
	}

	author.AuthorID, err = db.GetAuthorIDByName(author.AuthorName)
	if err == sql.ErrNoRows {
		log.Printf("-------NO RECORD------")
		author.AuthorID = db.InsertAuthor(author.AuthorName)
		book.BookAuthorID = author.AuthorID
	} else if err != nil {
		log.Fatal(err)
	} else {
		book.BookAuthorID = author.AuthorID
	}

	category := model.Category{
		CategoryName: r.FormValue("bookcategory"),
	}

	category.CategoryID, err = db.GetCategoryIDByName(category.CategoryName)
	if err == sql.ErrNoRows {
		log.Printf("-------NO RECORD------")
		category.CategoryID = db.InsertCategory(category.CategoryName)
		book.BookCategoryID = category.CategoryID
	} else if err != nil {
		log.Fatal(err)
	} else {
		book.BookCategoryID = category.CategoryID
	}

	publisher := model.Publisher{
		PublisherName: r.FormValue("bookpublisher"),
	}

	publisher.PublisherID, err = db.GetPublisherIDByName(publisher.PublisherName)
	if err == sql.ErrNoRows {
		log.Printf("-------NO RECORD------")
		publisher.PublisherID = db.InsertPublisher(publisher.PublisherName)
		book.BookPublisherID = publisher.PublisherID
	} else if err != nil {
		log.Fatal(err)
	} else {
		book.BookPublisherID = publisher.PublisherID
	}

	fmt.Println(book)

	err = db.InsertBook(book)
	if err != nil {
		log.Fatal(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		jsonbook, _ := json.Marshal(book)

		fmt.Fprintf(w, string(jsonbook))
	}

	bookid, err := db.GetBookIDByName(book.BookName)
	tag := r.FormValue("tags")
	taglist := strings.Split(tag, " ")

	for _, j := range taglist {
		err := db.InsertTag(j, bookid)
		db.CheckErr(err)
	}

	booklist := strings.Split(book.BookName, " ")

	for _, j := range booklist {
		err := db.InsertTag(j, bookid)
		db.CheckErr(err)
	}

	authorlist := strings.Split(author.AuthorName, " ")

	for _, j := range authorlist {
		err := db.InsertTag(j, bookid)
		db.CheckErr(err)
	}

	err3 := db.InsertTag(category.CategoryName, bookid)
	db.CheckErr(err3)
}

func GiveBook(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "giveBook.html", nil)
}

func GiveBookToUser(w http.ResponseWriter, r *http.Request) {
	NotPost(w, r)

	bookName := r.FormValue("bookname")
	userName := r.FormValue("username")

	bookID, err1 := db.GetBookIDByName(bookName)
	userID, err2 := db.GetIDByUsername(userName)

	w.Header().Set("Content-Type", "application/json")

	if err1 == sql.ErrNoRows {
		str := "ERROR NO RECORD BOOK"
		jsonstr, _ := json.Marshal(str)
		fmt.Fprintf(w, string(jsonstr))
	} else if err2 == sql.ErrNoRows {
		str := "ERROR NO RECORD USER"
		jsonstr, _ := json.Marshal(str)
		fmt.Fprintf(w, string(jsonstr))
	} else if err1 != nil || err2 != nil {
		str := "SOMETHINGS GOING WRONG"
		jsonstr, _ := json.Marshal(str)
		fmt.Fprintf(w, string(jsonstr))
	} else {
		err3 := db.InsertBooksAndUsers(bookID, userID)

		if err3 != nil {
			str := "SOMETHINGS GOING WRONG"
			jsonstr, _ := json.Marshal(str)
			fmt.Fprintf(w, string(jsonstr))

		} else {
			str := "PROCESS IS SUCCESSFUL"
			jsonstr, _ := json.Marshal(str)
			fmt.Fprintf(w, string(jsonstr))
		}
	}
}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "searchBook.html", nil)
}

func SearchBookAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	search := r.FormValue("search")
	search = strings.ToUpper(search)

	searchlist := strings.Split(search, " ")

	var bookidlist []int
	for _, j := range searchlist {
		bkid := db.GetBookIDByTag(j)
		bookidlist = append(bookidlist, bkid)
	}

	fmt.Println(bookidlist)

	uniqueidlist := model.UniqueSlice(bookidlist)
	fmt.Println(uniqueidlist)

	var booklist []model.Book

	for _, j := range uniqueidlist {
		bk := db.GetBookAllByID(j)
		booklist = append(booklist, bk)
	}

	jsonbooklist, _ := json.Marshal(booklist)
	fmt.Fprintf(w, string(jsonbooklist))
}
