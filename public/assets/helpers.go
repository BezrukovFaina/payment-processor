package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

func ParseJSON(body []byte, target interface{}) error {
	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("failed to parse json: %w", err)
	}
	return nil
}

func HandleError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("error: %v\n", err)
	}
}

func ValidateRequest(r *http.Request) error {
	if r == nil {
		return errors.New("request is nil")
	}
	if r.Method != http.MethodPost {
		return errors.New("only POST requests are allowed")
	}
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("invalid content type")
	}
	return nil
}