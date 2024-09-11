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
	//Test Eager Loading Staff
	// var staff models.Staff
	// facades.Orm().Query().With("Store.Manager").With("Address").Find(&staff)

	//Test Film
	var film models.Film
	facades.Orm().Query().Where(models.Film{ID: 83}).Find(&film)

	return ctx.Response().Success().Json(http.Json{
		// "data": staff,
		"data": film,
	})
}
