package auth

import (
	"backend-go/internal/model/form"
	"backend-go/internal/model/migrate"
	"backend-go/internal/model/response"
	"backend-go/internal/repository/users"
	"backend-go/pkg/token"
	"fmt"
	"time"
)

func Register(req form.RegisterForm) (*migrate.User, error) {
	hashed, err := token.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	verify, _ := users.GetUserByEmail(req.Email)
	if verify.Email == req.Email {
		return nil, err
	}

	user := migrate.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashed,
		CreatedAt: time.Now(),
	}

	createdUser, err := users.CreateUsers(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func Login(req form.LoginForm) (*response.Login, error) {
	users, err := users.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("email or password is incorrect")
	}

	if !token.CheckPasswordHash(req.Password, users.Password) {
		return nil, fmt.Errorf("email or password is incorrect")
	}

	jwtToken, err := token.GenerateToken(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	res := response.Login{
		Token: jwtToken,
	}
	return &res, nil
}
