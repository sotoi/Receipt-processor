package models

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required,shortDescription"`
	Price string `json:"price" validate:"required,price"`
}