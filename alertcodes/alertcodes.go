package alertcodes

type AlertMessage struct {
	UserID     int     `json:"userid"`
	ActionType string  `json:"action"`
	Amount     float32 `json:"amount"`
	Time       int     `json:"time"`
}

type AlertResponse struct {
	IsAlert    bool  `json:"alert"`
	AlertCodes []int `json:"alert_codes,omitempty"`
	UserID     int   `json:"user_id"`
}

const (
	DEPOSIT    = "deposit"
	WITHDRAWAL = "withdrawal"
)

func QueryAlertCodes(alertMessage AlertMessage) (AlertResponse, error) {

	var alertResponse AlertResponse
	alertResponse.UserID = alertMessage.UserID

	if alertMessage.ActionType == DEPOSIT {
		alertResponse.AlertCodes, alertResponse.IsAlert = checkDepositHistory(alertMessage)
		return alertResponse, nil
	}

	alertResponse.AlertCodes, alertResponse.IsAlert = checkWithdrawalHistory(alertMessage)
	return AlertResponse{}, nil
}

func checkDepositHistory(alertMessage AlertMessage) ([]int, bool) {
	return []int{}, false
}

func checkWithdrawalHistory(alertMessage AlertMessage) ([]int, bool) {
	return []int{}, false
}
