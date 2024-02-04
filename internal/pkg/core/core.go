package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Api struct {
	url   string
	token string
}

// TODO: good http request error handling

func NewCore() *Api {
	pwd := os.Getenv("api_pwd")
	url := os.Getenv("api_url")

	body, err := json.Marshal(map[string]string{
		"pwd": pwd,
	})
	if err != nil {
		panic(err)
	}
	res, err := http.Post(url+"/discord/bot-login", "application/json", bytes.NewBuffer(body))
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

	log.Println("Logged in to API Core successfully")
	return &Api{
		url:   url,
		token: result.Token,
	}
}
