package auth

import (
	"context"
	"errors"
	"fmt"
	"main/internal/services/auth"
)

type UseCase struct {
	auth  Auth
	user  User
	cache Cache
	email Email
}

func NewUseCase(auth Auth, user User, cache Cache, email Email) *UseCase {
	return &UseCase{auth: auth, user: user, cache: cache, email: email}
}

func (au UseCase) SignIn(ctx context.Context, data auth.SignIn) (string, error) {
	userDetail, err := au.user.GetByEmail(ctx, data.Email)
	if err != nil {
		return "", errors.New("user not found")
	}
	isValidPassword := au.auth.CheckPasswordHash(data.Password, *userDetail.Password)
	if !isValidPassword {
		return "", errors.New("invalid password")
	}
	var generateTokenData auth.GenerateToken
	generateTokenData.Email = userDetail.Email
	generateTokenData.Role = userDetail.Role
	token, err := au.auth.GenerateToken(ctx, generateTokenData)
	return token, err
}

func (au UseCase) SignUp(ctx context.Context, data auth.SignUp) (string, error) {
	hashPassword, err := au.auth.HashPassword(data.Password)
	if err != nil {
		return "", err
	}
	_, err = au.user.GetByEmail(ctx, data.Email)
	if err == nil {
		return "", errors.New("email already taken")
	}
	data.Password = hashPassword
	detail, err := au.user.Create(ctx, data)
	if err != nil {
		return "", err
	}
	var generateTokenData auth.GenerateToken
	generateTokenData.Email = detail.Email
	generateTokenData.Role = detail.Role
	token, err := au.auth.GenerateToken(ctx, generateTokenData)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (au UseCase) ForgotPsw(ctx context.Context, email string) (string, error) {
	_, err := au.user.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("email not found")
	}

	token, err := au.auth.GenerateResetToken(16)
	if err != nil {
		return "", err
	}

	code := au.email.GenerateCode(6)
	err = au.email.SendMailSimple("Password Reset Code", "Your reset code is: "+code, []string{email})
	if err != nil {
		return "", err
	}

	data := auth.ResetData{Email: email, Code: code}
	if err := au.cache.Set(ctx, token, data); err != nil {
		return "", err
	}
	return token, nil
}

func (au UseCase) CheckCode(ctx context.Context, code, token string) error {
	var data auth.ResetData
	if err := au.cache.Get(ctx, token, &data); err != nil {
		return err
	}
	fmt.Println("data", data.Code != code)
	if code != data.Code {
		return errors.New("code error")
	}
	return nil
}

func (au UseCase) UpdatePsw(ctx context.Context, data auth.UpdatePsw) error {
	var resetData auth.ResetData

	if err := au.cache.Get(ctx, data.Token, &resetData); err != nil {
		return err
	}

	hashPassword, err := au.auth.HashPassword(data.Password)
	if err != nil {
		return err
	}
	if err = au.user.UpdatePassword(ctx, resetData.Email, hashPassword); err != nil {
		return err
	}
	if err = au.cache.Delete(ctx, data.Token); err != nil {
		return err
	}
	return nil
}

func (au UseCase) R() {

}
