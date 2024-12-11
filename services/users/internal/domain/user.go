package domain

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     Email     `db:"email"`
	Password  Password  `db:"password"`
}

func NewUser(firstName, lastName, email, password string) (*User, error) {
	e := Email(email)
	if err := e.Validate(); err != nil {
		return nil, err
	}
	p, err := NewPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     e,
		Password:  p,
	}, nil
}
