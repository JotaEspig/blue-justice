package handlers

import (
	"net/http"
	auth "sigma/services/authentication"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	errNoResult = "sql: no rows in result set"
)

// Gets the token and sends a JSON containing information about the user to the browser
// if the token is valid
func GetLoggedUserInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Defines a function that sends an unauthorized code to the server
		unauthorizedJSON := func(json gin.H) {
			ctx.JSON(
				http.StatusUnauthorized,
				json,
			)
		}

		token, err := ctx.Cookie("auth")
		if token == "" || err != nil {
			unauthorizedJSON(nil)
			return
		}

		//dToken means decoded token
		dToken, err := defaultJWT.ValidateToken(token)
		if err != nil || !dToken.Valid {
			unauthorizedJSON(nil)
			return
		}

		claims := dToken.Claims.(jwt.MapClaims)

		now := time.Now().Unix()
		expiresAt := claims["exp"].(float64)
		if float64(now) > expiresAt {
			unauthorizedJSON(nil)
			return
		}

		user, err := auth.GetUser(db, claims["username"].(string))
		if err != nil {
			if err.Error() == errNoResult {
				ctx.Status(http.StatusUnauthorized)
				return
			}
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(
			http.StatusOK,
			gin.H{
				"claims": claims,
				"user":   user.ToMap(),
			},
		)
	}
}
