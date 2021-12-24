package main

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	users []*User
}

func NewUserRepo() *UserRepo {
	p1, _ := bcrypt.GenerateFromPassword([]byte("1111111111"), bcrypt.DefaultCost)
	p2, _ := bcrypt.GenerateFromPassword([]byte("2222222222"), bcrypt.DefaultCost)
	users := []*User{
		&User{
			ID:       1,
			Email:    "test@mail.com",
			Name:     "Name1",
			Password: string(p1),
		},
		&User{
			ID:       2,
			Email:    "test2@mail.com",
			Name:     "Name2",
			Password: string(p2),
		},
	}
	return &UserRepo{users: users}
}

func (ur *UserRepo) GetByEmail(email string) (*User, error) {
	for _, user := range ur.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (ur *UserRepo) GetByID(id int) (*User, error) {
	for _, user := range ur.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
