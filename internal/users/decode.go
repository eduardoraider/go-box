package users

import (
	"encoding/json"
	"io"
)

func decode(body io.ReadCloser) (*User, error) {
	defer body.Close()

	u := new(User)

	err := json.NewDecoder(body).Decode(u)
	if err != nil {
		return nil, err
	}

	return u, nil

}

func DecodeAndCreate(body io.ReadCloser) (*User, error) {
	defer body.Close()

	u, err := decode(body)
	err = encryptPassword(u)
	if err != nil {
		return nil, err
	}

	err = u.Validate()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func DecodeAndModify(body io.ReadCloser, u *User) (*User, error) {
	defer body.Close()
	mu, err := decode(body)
	if err != nil {
		return nil, err
	}

	err = u.ChangeName(mu.Name)
	if err != nil {
		return nil, err
	}

	return mu, nil
}
