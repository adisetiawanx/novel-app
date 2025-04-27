package middleware

import (
	"github.com/adisetiawanx/novel-app/internal/app"
	"github.com/adisetiawanx/novel-app/internal/helper"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func IsAdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookie, err := ctx.Cookie("token")
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing access token")
			}

			token, err := jwt.ParseWithClaims(cookie.Value, &helper.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.Config.Token.AccessSecret), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}

			claims, ok := token.Claims.(*helper.CustomClaims)
			if !ok || claims.Role != "admin" {
				return echo.NewHTTPError(http.StatusForbidden, "Access denied")
			}

			ctx.Set("user_id", claims.UserID)
			ctx.Set("role", claims.Role)

			return next(ctx)
		}
	}
}
