package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"midnite.com/takehometest/alertcodes"
)

func AlertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only accepts POST requests", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	// responseBody ALWAYS needs reading to ensure the tcp read buffer is drained and prevent memory leak
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "err reading response body", http.StatusInternalServerError)
		return
	}

	var newAlert alertcodes.AlertMessage
	err = json.Unmarshal(reqBody, &newAlert)
	if err != nil {
		http.Error(w, "failed to unmarshal json", http.StatusInternalServerError)
		return
	}

	err = validateAlertBody(newAlert)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := alertcodes.QueryAlertCodes(newAlert)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(respJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func validateAlertBody(alertBody alertcodes.AlertMessage) error {
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
