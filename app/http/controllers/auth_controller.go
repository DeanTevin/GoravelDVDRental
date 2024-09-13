package controllers

import (
	"goravel/app/helpers"
	"goravel/app/http/requests"
	"goravel/app/models"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthController struct {
	//Dependent services
	validate *helpers.RequestValidationHelper
}

func NewAuthController() *AuthController {
	return &AuthController{
		//Inject services
	}
}

func (r *AuthController) OauthLogin(ctx http.Context) http.Response {
	var RequestValidation requests.LoginAuthRequest
	errorResponseValidation := r.validate.ValidateRequest(&RequestValidation, ctx)
	if errorResponseValidation != nil {
		return ctx.Response().Status(406).Json(errorResponseValidation)
	}

	var staff models.Staff
	err := facades.Orm().Query().Where(models.Staff{Username: RequestValidation.Username}).FirstOrFail(&staff)
	if err != nil {
		return ctx.Response().Status(404).Json(http.Json{
			"error": "No Username Found for User",
		})
	}

	if !facades.Hash().Check(RequestValidation.Password, staff.Password) {
		return ctx.Response().Status(401).Json(http.Json{
			"error": "Password is Invalid",
		})
	}

	token, _ := facades.Auth(ctx).Login(&staff)

	payload, err := facades.Auth(ctx).Parse(token)

	if err != nil {
		return ctx.Response().Status(500).Json(http.Json{
			"error": "token formed incorrectly",
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"data": map[string]any{
			"token":      token,
			"issued_at":  payload.IssuedAt,
			"expired_at": payload.ExpireAt,
		},
	})
}

func (r *AuthController) RefreshToken(ctx http.Context) http.Response {
	_, err := facades.Auth(ctx).Parse(ctx.Request().Header("Authorization"))

	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error":  err.Error(),
			"status": "Invalid Token",
		})
	}

	token, err := facades.Auth(ctx).Refresh()

	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error":  err.Error(),
			"status": "Invalid Token",
		})
	}

	payload, err := facades.Auth(ctx).Parse(token)

	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error":  err.Error(),
			"status": "Invalid Token",
		})
	}

	return ctx.Response().Json(http.StatusAccepted, http.Json{
		"data": map[string]any{
			"token":        token,
			"expired_at": payload.ExpireAt.Format(time.RFC1123),
			"issued_at":  payload.ExpireAt.Format(time.RFC1123),
			"status":     "Refresh token success!",
		},
	})
}

func (r *AuthController) OauthLogout(ctx http.Context) http.Response {
	_, err := facades.Auth(ctx).Parse(ctx.Request().Header("Authorization"))

	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error":  err.Error(),
			"status": "Invalid Token",
		})
	}

	err = facades.Auth(ctx).Logout()

	if err != nil {
		ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error":  err.Error(),
			"status": "Invalid Token",
		})
	}

	return ctx.Response().Json(http.StatusAccepted, http.Json{
		"status": "Log out success!",
	})
}
