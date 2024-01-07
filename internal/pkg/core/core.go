package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Api struct {
	url   string
	token string
}

var Core *Api

func init() {
	pwd := os.Getenv("api_pwd")
	url := os.Getenv("api_url")

	body, err := json.Marshal(map[string]string{
		"pwd": pwd,
	})
	if err != nil {
		panic(err)
	}
	res, err := http.Post(url+"/discord/bot-login"+pwd, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var result struct {
		Token string `json:"token"`
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(resBody, &result)
	if err != nil {
		panic(err)
	}

	Core = &Api{
		url:   url,
		token: result.Token,
	}
}
