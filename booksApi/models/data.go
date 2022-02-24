package models

var db []Book

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}
type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		ID:            1,
		Title:         "The Lord of the Rings",
		YearPublished: 1954,
		Author: Author{
			Name:     "John Ronald",
			LastName: "Tolkien",
			BornYear: 1892,
		},
	}
	db = append(db, book1)
}
