package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type AlertMessage struct {
	UserID     int     `json:"userid"`
	ActionType string  `json:"action"`
	Amount     float32 `json:"amount"`
	Time       int     `json:"time"`
}

func AlertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only accepts POST requests", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	// responseBody ALWAYS needs reading to ensure the tcp read buffer is drained and prevent memory leak
	respBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "err reading response body", http.StatusInternalServerError)
		return
	}

	var newAlert AlertMessage
	err = json.Unmarshal(respBody, &newAlert)
	if err != nil {
		http.Error(w, "failed to unmarshal json", http.StatusInternalServerError)
		return
	}

	err = validateAlertBody(newAlert)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

}

func validateAlertBody(alertBody AlertMessage) error {
	if alertBody.UserID <= 0 {
		return errors.New("userID must be bigger than 0")
	}

	switch alertBody.ActionType {
	case "withdrawal", "deposit":
	default:
		return errors.New("not supported method")
	}

	if alertBody.Amount < 0 {
		return errors.New("amount is less than 0")
	}

	if alertBody.Time < 0 {
		return errors.New("time is less than 0")
	}

	return nil
}
