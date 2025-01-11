package router

import (
	"encoding/json"
	"errors"
	"fetch/models"
	"fetch/service"
	"fetch/store"
	"fetch/validator"
	"log"
	"net/http"

	"github.com/google/uuid"
)

var (
	notValidReceipt = "The receipt is invalid."
	notFoundReceipt = "No receipt found for that ID."
)

type Handler struct {
	receiptService *service.ReceiptService
	validate       *validator.Validate
}

func NewHandler(rs *service.ReceiptService, validator *validator.Validate) *Handler {
	return &Handler{
		receiptService: rs,
		validate:       validator,
	}
}

func (h *Handler) SaveReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var receipt *models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, notValidReceipt, http.StatusBadRequest)
		return
	}

	err = h.validate.Validate(receipt)
	if err != nil {
		log.Printf("validation error: %v", err)
		http.Error(w, notValidReceipt, http.StatusBadRequest)
		return
	}
	h.receiptService.SaveReceipt(receipt)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": receipt.ID})

}

func (h *Handler) GetPoints(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("id")

	_, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "ID format not valid", http.StatusBadRequest)
		return
	}

	points, err := h.receiptService.GetPoints(id)
	if errors.Is(err, store.ErrNotExists) {
		http.Error(w, notFoundReceipt, http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
