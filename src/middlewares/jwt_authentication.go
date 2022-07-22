package middlewares

import (
	um "crowdfunding-api/src/modules/user"
	"crowdfunding-api/src/utils"
	"crowdfunding-api/src/utils/common"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JwtAuthMiddleware Token Authentication
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Extract Token Data
		tokenString := utils.ExtractToken(context)

		// Check token string
		if tokenString == "" {
			// Return Error
			common.SendError(context, http.StatusUnauthorized, "Unauthorized", []string{"Authentication Token Required"})

			// Abort to do next handler
			context.Abort()
			return
		}

		// Get Secret Key from ENV
		key := os.Getenv("API_JWT_SECRET")

		// Parse JWT and validate
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check sign in method token
			if jwt.GetSigningMethod("HS256") != token.Method {
				// When sign in method not same
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// return key
			return []byte(key), nil
		})

		// Check if user exist in database & Token Expired
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Catch User Email
			email := fmt.Sprintf("%v", claims["email"])

			// Check if token expired
			if time.Now().Unix() > int64(math.Round(claims["expired"].(float64))) {
				// Return Error
				common.SendError(context, http.StatusUnauthorized, "Unauthorized", []string{"Token expired"})

				// Abort to do next handler
				context.Abort()
				return
			}

			repository := um.GetRepository()

			// Check User On Database
			_, err = repository.FindByEmail(email)

			if err != nil {
				// Return Error
				common.SendError(context, http.StatusUnauthorized, "Unauthorized", []string{"User not found"})

				// Abort to do next handler
				context.Abort()
				return
			}

		} else {
			// Token not valid
			common.SendError(context, http.StatusUnauthorized, "Unauthorized 1", []string{err.Error()})

			// Abort to do next handler
			context.Abort()
			return
		}

		// Check error
		if token == nil && err != nil {
			// Token not valid
			common.SendError(context, http.StatusUnauthorized, "Unauthorized 2", []string{err.Error()})

			// Abort to do next handler
			context.Abort()
			return
		}

		// Next
		context.Next()
	}
}
