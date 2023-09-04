package middlewares

import (
	"fmt"
	"gologin/config"
	"gologin/inits"
	"gologin/response"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("RequireAuth")
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.Abort()
			response.FailWithCode(response.ResponseCodeUnauthorized, ctx)
			return
		}
		tokenString := strings.TrimPrefix(token, "Bearer ")
		fmt.Println(tokenString)
		logrus.Info("token", tokenString)
		tokenClaims, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Env.SECRET), nil
		})
		if err != nil {
			logrus.Error("token pare error")
			response.FailWithCode(response.ResponseCodeUnauthorized, ctx)
			ctx.Abort()
			return
		}
		if !tokenClaims.Valid {
			ctx.Abort()
			logrus.Error("token invalid")
			response.FailWithCode(response.ResponseCodeUnauthorized, ctx)
			return
		}
		if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				response.FailWithCode(response.ResponseCodeUnauthorized, ctx)
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			email, ok := claims["email"].(string)
			if ok {
				// 通过key获取redis
				redisToken := inits.RedisClient.Get(email).Val()
				if tokenString != redisToken {
					response.FailWithCode(response.ResponseCodeUnauthorized, ctx)
					ctx.Abort()
					return
				}
				ctx.Set("email", email)
			} else {
				response.FailServer(ctx)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}

}
