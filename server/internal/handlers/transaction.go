package handlers

import (
	"encoding/json"
	"net/http"

	"para-bank/server/internal/models"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	// Логіка для отримання транзакцій
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Transaction{})
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// Логіка для створення транзакції
	w.WriteHeader(http.StatusCreated)
}
