package repository

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type BookRepository struct {
	Resources []Book
}

func (br BookRepository) Finds(ids []int) []Book {
	var result []Book

	for _, id := range ids {
		for _, book := range br.Resources {
			if book.ID == id {
				result = append(result, book)
			}
		}
	}

	return result
}

func (br BookRepository) FindAll() []Book {
	return br.Resources
}

var BookResource = BookRepository{
	Resources: []Book{
		Book{
			ID:    1,
			Title: "Test One",
			Year:  2022,
		},
		Book{
			ID:    2,
			Title: "Test Two",
			Year:  2022,
		},
		Book{
			ID:    3,
			Title: "Test Three",
			Year:  2022,
		},
		Book{
			ID:    4,
			Title: "Test Four",
			Year:  2022,
		},
		Book{
			ID:    5,
			Title: "Test Five",
			Year:  2022,
		},
		Book{
			ID:    6,
			Title: "Test Six",
			Year:  2022,
		},
	},
}
