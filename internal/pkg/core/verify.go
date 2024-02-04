package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nobypass/fds-bot/internal/pkg/models"
	"net/http"
)

func (a *Api) Verify(id, name, nick string) (*models.DiscordVerifyResponse, error) {
	body := models.DiscordVerifyInput{
		Name: name,
		Nick: nick,
		ID:   id,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, a.url+"/discord/verify", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+a.token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	} else if res.StatusCode != http.StatusOK {
		msg := attemptDecodeHttpErr(res.Body)
		return nil, fmt.Errorf("%d %s", res.StatusCode, msg)
	}

	var resp models.DiscordVerifyResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, err
	} else if resp.Actual == "" {
		return nil, fmt.Errorf("empty response")
	}

	return &resp, nil
}
