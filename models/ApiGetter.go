package models

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Region string

const (
	Americas Region = "americas"
	Asia     Region = "asia"
	Europe   Region = "europe"
	SEA      Region = "sea"
)

type ApiGetter struct {
	Client *http.Client
	ApiKey string
	Region Region
}

func NewApiClient(api_key string, region Region) (ApiGetter, error) {
	Client := &http.Client{}
	return ApiGetter{Client, api_key, region}, nil
}

func (g ApiGetter) Get(url string) string {
	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s.api.riotgames.com%s", g.Region, url), nil)
	if err != nil {
		fmt.Println("Error creating request")
		log.Print(err)
		return ""
	}
	r.Header.Add("X-Riot-Token", g.ApiKey)
	resp, err := g.Client.Do(r)
	if err != nil {
		fmt.Println("Error sending request")
		log.Print(err)
		return ""
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response")
		log.Print(err)
		return ""
	}
	return string(resBody)
}

func (g ApiGetter) Getf(url string, args ...interface{}) string {
	return g.Get(fmt.Sprintf(url, args...))
}
