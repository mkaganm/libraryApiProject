package model

type User struct {
	ID        int
	UserName  string // `json:username`
	Password  string // `json:password`
	FirstName string // `json:firstname`
	LastName  string // `json:lastname`
	Email     string // `json:email`
	Phone     string // `json:phone`
}

type Admin struct {
	ID        int
	AdminName string
	Password  string
}

type Book struct {
	BookID          int
	BookName        string
	BookAuthorID    int
	BookCategoryID  int
	BookPublisherID int
	BookAmount      string
}

type Author struct {
	AuthorID   int
	AuthorName string
}

type Category struct {
	CategoryID   int
	CategoryName string
}

type Publisher struct {
	PublisherID   int
	PublisherName string
}

type Tag struct {
	Tagid  int
	Tag    string
	BookID int
}

type MKM struct {
	ID int
}

// * ONLY UNIQUE ELEMENST FOR SLICE
func UniqueSlice(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
