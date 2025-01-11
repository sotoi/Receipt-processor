package service

import (
	"fetch/models"
	"fetch/rules"
	"fetch/store"
)

// ReceiptService is a struct that contains a storage and a rule engine calculating receipt points
type ReceiptService struct {
	storage store.Storage
	rules	 rules.RuleRunner
}

// NewReceiptService is a constructor for ReceiptService
func NewReceiptService(storage store.Storage, rules rules.RuleRunner) *ReceiptService {
	return &ReceiptService{
		storage: storage,
		rules: rules,
	}
}

//	SaveReceipt saves a receipt to the storage and returns the receipt ID
func (rs *ReceiptService) SaveReceipt(receipt *models.Receipt) string {
	rs.storage.SaveReceipt(receipt)
	return receipt.ID
}

//	GetPoints calculates the points for a receipt and returns the points
func (rs *ReceiptService) GetPoints(id string) (int, error) {
	receipt, err := rs.storage.GetReceipt(id)
	if err != nil {
		return 0, err
	}
	points := rs.rules.CalculatePoints(receipt)
	return points, nil
}