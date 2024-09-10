package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	var staff models.Staff
	facades.Orm().Query().With("Store.Manager").With("Address").Find(&staff)
	// facades.Orm().Query().Find(&staff)

	return ctx.Response().Success().Json(http.Json{
		"data": staff,
	})
}
