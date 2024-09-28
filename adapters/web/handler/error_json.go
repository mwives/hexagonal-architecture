package handler

import "encoding/json"

func jsonError(msg string) []byte {
	errorMessage := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, err := json.Marshal(errorMessage)
	if err != nil {
		return []byte(`{"error": "internal server error"}`)
	}
	return r
}
