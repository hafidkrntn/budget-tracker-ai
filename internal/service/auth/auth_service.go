package auth

import (
	"backend-go/config"
	"backend-go/internal/model/form"
	"backend-go/internal/model/migrate"
	"backend-go/internal/model/response"
	"backend-go/internal/repository/users"
	"backend-go/pkg/token"
	"fmt"
	"time"
)

func Register(req form.RegisterForm) (*migrate.User, error) {
	db := config.GetDB()

	hashed, err := token.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	verify, err := users.GetUserByEmail(db, req.Email)
	if err == nil && verify.Email == req.Email {
		return nil, fmt.Errorf("email already in use")
	}

	user := migrate.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashed,
		CreatedAt: time.Now(),
	}

	createdUser, err := users.CreateUsers(db, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func Login(req form.LoginForm) (*response.Login, error) {
	db := config.GetDB()

	user, err := users.GetUserByEmail(db, req.Email)
	if err != nil {
		return nil, fmt.Errorf("email or password is incorrect")
	}

	if !token.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("email or password is incorrect")
	}

	jwtToken, err := token.GenerateToken(user.Email, user.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	res := response.Login{
		Token: jwtToken,
	}
	return &res, nil
}
