package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type LoginAuthRequest struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (r *LoginAuthRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *LoginAuthRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"username": "required|string|max_len:16",
		"password": "required|string|min_len:8",
	}
}

func (r *LoginAuthRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *LoginAuthRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *LoginAuthRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
