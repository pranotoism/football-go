package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotoism/football-go/util"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				util.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
