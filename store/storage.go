package store

import "fetch/models"

// Storage is an interface that defines the SaveReceipt and GetReceipt methods can be implemented by any storage example Postgres, MySQL, InMemory
type Storage interface {
	// SaveReceipt saves a receipt to the storage
	SaveReceipt(*models.Receipt)
	// GetReceipt returns a receipt from the storage
	GetReceipt(string) (*models.Receipt, error)
}