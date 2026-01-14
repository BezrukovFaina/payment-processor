package paymentprocessor

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

// generateToken generates a unique token
func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// hashPassword hashes a password
func hashPassword(password string) (string, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(password)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// verifyPassword verifies a password against a hash
func verifyPassword(password string, hash string) (bool, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return false, err
	}
	return hashedPassword == hash, nil
}

// currentTime returns the current time
func currentTime() time.Time {
	return time.Now().UTC()
}

// validateTransaction validates a transaction
func validateTransaction(amount float64, currency string) error {
	if amount <= 0 {
		return errors.New("invalid transaction amount")
	}
	if len(currency) != 3 {
		return errors.New("invalid currency code")
	}
	return nil
}