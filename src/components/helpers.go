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

// PaymentRequest represents a payment request
type PaymentRequest struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	Timestamp time.Time `json:"timestamp"`
}

// ValidatePaymentRequest checks if a payment request is valid
func ValidatePaymentRequest(req *PaymentRequest) error {
	if req.Amount <= 0 {
		return errors.New("payment amount must be greater than zero")
	}
	if req.Currency == "" {
		return errors.New("payment currency is required")
	}
	if req.Timestamp.IsZero() {
		return errors.New("payment timestamp is required")
	}
	return nil
}

// GeneratePaymentID generates a unique payment ID
func GeneratePaymentID() string {
	return uuid.New().String()
}

// ParsePaymentRequestFromJSON parses a payment request from JSON
func ParsePaymentRequestFromJSON(data []byte) (*PaymentRequest, error) {
	var req PaymentRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil, err
	}
	req.ID = GeneratePaymentID()
	req.Timestamp = time.Now()
	return &req, nil
}

// HandlePaymentRequest handles a payment request
func HandlePaymentRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	data, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil || data <= 0 {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	body := make([]byte, data)
	_, err = r.Body.Read(body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusInternalServerError)
		return
	}
	req, err := ParsePaymentRequestFromJSON(body)
	if err != nil {
		http.Error(w, "failed to parse payment request", http.StatusBadRequest)
		return
	}
	err = ValidatePaymentRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("payment request received: %+v\n", req)
	w.WriteHeader(http.StatusCreated)
}

// RunServer starts the payment processor server
func RunServer(port int) error {
	http.HandleFunc("/payment", HandlePaymentRequest)
	return http.ListenAndServe(":"+strconv.Itoa(port), nil)
}