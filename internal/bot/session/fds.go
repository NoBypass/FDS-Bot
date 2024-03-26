package session

import (
	"bytes"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"os"
)

type FDSConnection struct {
	url   string
	token string
}

func ConnectToFDS(tracer opentracing.Tracer) *FDSConnection {
	conn := &FDSConnection{
		url: os.Getenv("API_URL"),
	}

	sp := tracer.StartSpan("Logging bot in")
	defer sp.Finish()
	resp, err := conn.Login(sp, os.Getenv("PASSWORD"))
	if err != nil {
		return nil
	}

	conn.token = resp.Token
	return conn
}

func (c *FDSConnection) newRequest(method, path string, body interface{}, sp opentracing.Span) (*http.Request, error) {
	var data []byte
	if body != nil {
		d, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		data = d
	}
	reader := bytes.NewReader(data)
	req, err := http.NewRequest(method, c.url+path, reader)
	if err != nil {
		return nil, err
	}

	if sp != nil {
		err := sp.Tracer().Inject(
			sp.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(req.Header),
		)
		if err != nil {
			panic(err)
		}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	return req, nil
}

func do[T any](req *http.Request) (*T, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, err
	}

	v := new(T)
	err = json.NewDecoder(res.Body).Decode(v)
	return v, err
}
