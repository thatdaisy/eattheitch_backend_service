package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SessionSetup(router *gin.Engine) {

	store := cookie.NewStore([]byte("super-secret-key"))

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   6000, // 10 minutes
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	router.Use(sessions.Sessions("mysession", store))
}

func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)

		email := session.Get("email")
		if email == nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authentication required",
			})
			return
		}

		context.Set("email", email)

		context.Next()
	}
}
