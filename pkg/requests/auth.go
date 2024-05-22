package requests

import (
	"bytes"
	"context"
	"encoding/json"
	pb "github.com/eduardoraider/go-box/proto/v1/auth"
	"io"
	"os"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HTTPAuth(path, user, pass string) error {
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

	token, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return createTokenCache(string(token))
}

func GRPCAuth(user, pass string) error {
	creds := &pb.Credentials{
		Username: user,
		Password: pass,
	}

	conn := GetGRPCConn()
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	res, err := client.Login(context.Background(), creds)
	if err != nil {
		return err
	}

	return createTokenCache(res.Token)
}

type cacheToken struct {
	Token string `json:"token"`
}

func createTokenCache(token string) error {

	file, err := os.Create("./cache/.cacheToken")
	if err != nil {
		return err
	}
	defer file.Close()

	cache := cacheToken{token}
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
