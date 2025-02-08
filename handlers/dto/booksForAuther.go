package dto

type BooksForAutherDTO struct {
	AuthorName    []string `json:"author_name"`   
	Title         string   `json:"title"`
}
