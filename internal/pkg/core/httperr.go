package core

import (
	"encoding/json"
	"io"
)

func attemptDecodeHttpErr(body io.ReadCloser) string {
	var resp struct {
		Message string `json:"message"`
	}

	err := json.NewDecoder(body).Decode(&resp)
	if err != nil {
		return "unknown error"
	}

	return resp.Message
}
