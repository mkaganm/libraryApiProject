package main

import (
	"db"
	"handler"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var Tmpl *template.Template

func init() {
	Tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index)
	r.HandleFunc("/userLogin", handler.UserLoginPage)
	r.HandleFunc("/userRegister", handler.UserRegisterPage)
	r.HandleFunc("/adminLogin", handler.AdminLoginPage)
	r.HandleFunc("/userAuth", handler.UserAuth)
	r.HandleFunc("/adminAuth", handler.AdminAuth)
	r.HandleFunc("/registerAuth", handler.RegisterAuth)
	r.HandleFunc("/addBook", handler.AddBook)
	r.HandleFunc("/bookAuth", handler.BookAuth)
	r.HandleFunc("/bookList", handler.BookList)
	r.HandleFunc("/giveBook", handler.GiveBook)
	r.HandleFunc("/giveBookToUser", handler.GiveBookToUser)
	r.HandleFunc("/searchBook", handler.SearchBook)
	r.HandleFunc("/searchBookAuth", handler.SearchBookAuth)

	defer http.ListenAndServe(":8080", r)


}
