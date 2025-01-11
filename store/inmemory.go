package store

import (
	"errors"
	"fetch/models"
	"sync"

	"github.com/google/uuid"
)

// ErrNotExists is an error that is returned when a receipt does not exist
var ErrNotExists = errors.New("receipt does not exist")

// InMemory is a struct that contains a map of receipts in memory
type InMemory struct {
	receipts map[string]models.Receipt
	lock     sync.RWMutex
}

// NewInMemory is a constructor for InMemory
func NewInMemory() *InMemory {
	receipts := map[string]models.Receipt{}
	return &InMemory{
		receipts: receipts,
	}
}

// SaveReceipt saves a receipt to the in-memory storage
func (i *InMemory) SaveReceipt(receipt *models.Receipt) {
	i.lock.Lock()
	defer i.lock.Unlock()
	receipt.ID = uuid.NewString()
	i.receipts[receipt.ID] = *receipt
}

// GetReceipt returns a receipt from the in-memory storage
func (i *InMemory) GetReceipt(id string) (*models.Receipt, error) {
	i.lock.RLock()
	defer i.lock.RUnlock()

	receipt, exists := i.receipts[id]
	if !exists {
		return nil, ErrNotExists
	}
	return &receipt, nil
}
