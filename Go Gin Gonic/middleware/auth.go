package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {

		session := sessions.Default(c)

		sessionId := session.Get("id_user")

		if sessionId == nil {
			c.JSON(http.StatusNotFound, "Unauthorized")
			c.Abort()
		}
	}
}
