package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"goravel/app/http/middleware"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/staff", userController.Show)

	loginController := controllers.NewAuthController()
	facades.Route().Prefix("/oauth").Group(func(router route.Router) {
		router.Post("/login", loginController.OauthLogin)
	})

	facades.Route().Middleware(middleware.Auth()).Prefix("/oauth").Group(func(router route.Router) {
		router.Post("/refresh", loginController.RefreshToken)
		router.Post("/logout", loginController.OauthLogout)
	})
}
