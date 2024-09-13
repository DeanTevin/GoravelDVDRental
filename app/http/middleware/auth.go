package middleware

import (
	"errors"
	"goravel/app/models"
	"strconv"

	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Auth() http.Middleware {
	return func(ctx http.Context) {
		var account models.Staff
		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"Error":  "Unauthorized Access",
				"Status": "Unauthorized",
			})
			return
		}

		payload, err := facades.Auth(ctx).Parse(token)

		if err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
					"Error":  "JWT Token Expired",
					"Status": "Unathorized",
				})
				return
			}

			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"Error":  "JWT Token Invalid",
				"Status": "Unathorized",
			})
			return
		}

		staff_id, _ := strconv.Atoi(payload.Key)

		err = facades.Orm().Query().Where(models.Staff{ID: uint(staff_id)}).FirstOrFail(&account)

		if err != nil {
			ctx.Request().AbortWithStatusJson(http.StatusUnauthorized, http.Json{
				"Error":  "User Not Found or Invalid",
				"Status": http.StatusUnauthorized,
			})
			return
		}

		ctx.Request().Next()
	}
}
