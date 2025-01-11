package main

import (
	"fetch/router"
	"fetch/rules"
	"fetch/service"
	"fetch/store"
	"fetch/validator"
	"log"
	"net/http"
)

func main() {
	receiptService := initializeServices()

	mux :=initializeRouter(receiptService)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func initializeServices() *service.ReceiptService {
	inMemoryStorage := store.NewInMemory()
	ruleEngine := rules.NewRuleEngine(rules.DefaultRules)
	receiptService := service.NewReceiptService(inMemoryStorage, ruleEngine)
	return receiptService
}

func initializeRouter(receiptService *service.ReceiptService) *http.ServeMux {
	mux := http.NewServeMux()
	validate := validator.NewValidate()
	handler := router.NewHandler(receiptService, validate)

	mux.HandleFunc("/receipts/process", handler.SaveReceipt)
	mux.HandleFunc("/receipts/{id}/points", handler.GetPoints)

	return mux
}
