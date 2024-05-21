package requests

import (
	"errors"
	"io"
	"net/http"
)

func validateResponse(resp *http.Response) ([]byte, error) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 && resp.StatusCode < 600 {
		return nil, errors.New(string(data))
	}

	return data, nil
}

func Post(path string, body io.Reader) ([]byte, error) {
	resp, err := doRequest("POST", path, body, nil, false)
	if err != nil {
		return nil, err
	}

	return validateResponse(resp)
}

func AuthenticatedPostWithHeaders(path string, body io.Reader, headers map[string]string) ([]byte, error) {
	resp, err := doRequest("POST", path, body, headers, true)
	if err != nil {
		return nil, err
	}

	return validateResponse(resp)
}

func AuthenticatedPost(path string, body io.Reader) ([]byte, error) {
	resp, err := doRequest("POST", path, body, nil, true)
	if err != nil {
		return nil, err
	}

	return validateResponse(resp)
}

func AuthenticatedGet(path string) ([]byte, error) {
	resp, err := doRequest("GET", path, nil, nil, true)
	if err != nil {
		return nil, err
	}

	return validateResponse(resp)
}

func AuthenticatedPut(path string, body io.Reader) ([]byte, error) {
	resp, err := doRequest("PUT", path, body, nil, true)
	if err != nil {
		return nil, err
	}

	return validateResponse(resp)
}

func AuthenticatedDelete(path string) error {
	_, err := doRequest("DELETE", path, nil, nil, true)

	return err
}
