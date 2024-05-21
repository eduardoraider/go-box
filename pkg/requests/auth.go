package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(path, user, pass string) error {
	creds := &Credentials{user, pass}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(creds)
	if err != nil {
		return err
	}

	res, err := doRequest("POST", path, &body, nil, false)
	if err != nil {
		return err
	}

	return createTokenCache(res.Body)
}

type cacheToken struct {
	Token string `json:"token"`
}

func createTokenCache(body io.ReadCloser) error {
	token, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	file, err := os.Create("./cache/.cacheToken")
	if err != nil {
		return err
	}
	defer file.Close()

	cache := cacheToken{string(token)}
	data, err := json.Marshal(&cache)
	if err != nil {
		return err
	}

	_, err = file.Write(data)

	return err
}

func readCacheToken() (string, error) {
	data, err := os.ReadFile("./cache/.cacheToken")
	if err != nil {
		return "", err
	}

	var cache cacheToken

	err = json.Unmarshal(data, &cache)
	if err != nil {
		return "", err
	}

	return cache.Token, nil
}
