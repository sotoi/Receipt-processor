package models

type Receipt struct {
	ID           string
	Retailer     string `json:"retailer" validate:"required,retailer"`
	PurchaseDate string `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []Item `json:"items" validate:"required,gt=0,dive"`
	Total        string `json:"total" validate:"required,price"`
}
