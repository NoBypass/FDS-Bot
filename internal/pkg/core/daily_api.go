package core

import (
	"encoding/json"
	"github.com/nobypass/fds-bot/internal/pkg/models"
	"io/ioutil"
	"net/http"
)

func (a *Api) Daily(id string) (*models.DailyResponse, error) {
	req, err := http.NewRequest(http.MethodPatch, a.url+"/discord/"+id+"/daily", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+a.token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var member models.DailyResponse
	err = json.Unmarshal(body, &member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}
