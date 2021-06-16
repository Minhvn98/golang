package model

type Review struct {
	Id      int64  `json:"id"`
	BookId  int64  `json:"book_id"`
	Comment string `json:"comment"`
	Rating  int    `json:"rating"`
}
