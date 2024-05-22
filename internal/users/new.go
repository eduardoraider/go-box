package users

func New(id int64, name string, login string, password string) (*User, error) {
	u := &User{
		ID:       id,
		Name:     name,
		Login:    login,
		Password: password,
	}

	err := u.Validate()
	if err != nil {
		return nil, err
	}

	return u, nil
}
