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

	//Test Film (table with custom datatypes. Please see the Models in case u forget.)
	var film models.Film
	facades.Orm().Query().Where(models.Film{ID: 83}).With("Actors").With("Categories").
		With("Language").Find(&film)

	//this is for the Multiple row. See the [] in models.Film its create an array of data instead of single instance.
	// var film []models.Film
	// facades.Orm().Query().With("Actors").With("Categories").Find(&film)

	return ctx.Response().Success().Json(http.Json{
		// "data": staff,
		"data": film,
	})
}
