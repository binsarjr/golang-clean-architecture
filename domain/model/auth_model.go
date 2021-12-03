package model

import (
	"giapps/servisin/domain/vo"
	"giapps/servisin/infrastructure/exception"

	"net/http"
)

type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthLoginRequest) Bind(r *http.Request) error {
	if a.Username == "" {
		return exception.NewMissingRequired("username")
	}
	if a.Password == "" {
		return exception.NewMissingRequired("password")
	}

	return nil
}

type AuthRegisterRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (a *AuthRegisterRequest) Bind(r *http.Request) error {
	var err error
	if a.Username == "" {
		return exception.NewMissingRequired("username")
	}

	if a.Email == "" {
		return exception.NewMissingRequired("email")
	}
	_, err = vo.NewEmail(a.Email)
	if err != nil {
		return err
	}

	if a.Password == "" {
		return exception.NewMissingRequired("password")
	}
	if a.ConfirmPassword == "" {
		return exception.NewMissingRequired("konfirmasi password")
	}
	if a.ConfirmPassword != a.Password {
		return exception.NewValidationNotMatch("password", "konfirmasi password")
	}

	return nil
}

type AuthResponse struct {
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
}

func (u *AuthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
