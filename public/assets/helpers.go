package paymentprocessor

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// PaymentError represents an error that occurs during payment processing
type PaymentError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the error message
func (e *PaymentError) Error() string {
	return e.Message
}

// NewPaymentError returns a new PaymentError instance
func NewPaymentError(code int, message string) error {
	return &PaymentError{Code: code, Message: message}
}

// PaymentRequest represents a payment request
type PaymentRequest struct {
	ID        string  `json:"id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	CardNumber string `json:"card_number"`
	Expiry    string  `json:"expiry"`
	CVV       string  `json:"cvv"`
}

// Validate checks if the payment request is valid
func (pr *PaymentRequest) Validate() error {
	if pr.Amount <= 0 {
		return NewPaymentError(400, "invalid amount")
	}
	if len(pr.CardNumber) != 16 {
		return NewPaymentError(400, "invalid card number")
	}
	if len(pr.Expiry) != 5 {
		return NewPaymentError(400, "invalid expiry date")
	}
	if len(pr.CVV) != 3 {
		return NewPaymentError(400, "invalid cvv")
	}
	return nil
}

// ProcessPayment processes a payment request
func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var paymentRequest PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = paymentRequest.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paymentID := uuid.New().String()
	log.Printf("processing payment %s\n", paymentID)
	// simulate payment processing time
	time.Sleep(2 * time.Second)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"payment_id": paymentID})
}

// GetPaymentStatus returns the status of a payment
func GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	paymentID := r.URL.Query().Get("payment_id")
	if paymentID == "" {
		http.Error(w, "payment id is required", http.StatusBadRequest)
		return
	}
	// simulate database query to get payment status
	status := "success"
	if status == "success" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": status})
	} else {
		w.WriteHeader(http.StatusFailedDependency)
		json.NewEncoder(w).Encode(map[string]string{"status": status})
	}
}

// HandlePaymentError handles payment errors
func HandlePaymentError(w http.ResponseWriter, err error) {
	paymentError, ok := err.(*PaymentError)
	if !ok {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, paymentError.Message, paymentError.Code)
}