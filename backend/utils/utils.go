package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing Request Body in ParseJSON")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}
