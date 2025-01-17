package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PurchaseRequest struct {
	ProductValue uint64 `json:"productValue"`
}

type PurchaseResponse struct {
	Message        string `json:"message"`
	CashbackAmount uint64 `json:"cashbackAmount"`
	TokenPoints    uint64 `json:"tokenPoints"`
}

func BuyProductHandler(w http.ResponseWriter, r *http.Request) {
	var req PurchaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	cashbackAmount := (req.ProductValue * 10) / 100
	tokenPoints := uint64(10) 

	res := PurchaseResponse{
		Message:        "Purchase successful",
		CashbackAmount: cashbackAmount,
		TokenPoints:    tokenPoints,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetCashbackHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address parameter is required", http.StatusBadRequest)
		return
	}

	cashbackAmount := uint64(100)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"address":        address,
		"cashbackAmount": cashbackAmount,
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/purchase", BuyProductHandler).Methods("POST")
	r.HandleFunc("/cashback", GetCashbackHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}