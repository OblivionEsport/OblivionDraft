package models

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ImOlli/go-lcu/lcu"
)

type Getter struct {
	Client    *http.Client
	AuthToken string
	BaseURL   string
}

func NewGetter() (Getter, error) {
	info, err := lcu.FindLCUConnectInfo()

	if err != nil {
		if lcu.IsProcessNotFoundError(err) {
			return Getter{}, err
		}

		return Getter{}, err
	}

	log.Printf("LeagueClient is running on port %s and you can authenticate with following token: %s", info.Port, info.AuthToken)

	AuthToken := base64.StdEncoding.EncodeToString([]byte("riot:" + info.AuthToken))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // todo : add cert instead of skip
	}
	Client := &http.Client{Transport: tr}
	BaseURL := fmt.Sprintf("https://127.0.0.1:%s", info.Port)
	return Getter{Client, AuthToken, BaseURL}, nil
}

func (g Getter) Get(url string) string {

	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", g.BaseURL, url), nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", g.AuthToken))
	resp, err := g.Client.Do(r)
	if err != nil {
		panic(err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(resBody)
}
