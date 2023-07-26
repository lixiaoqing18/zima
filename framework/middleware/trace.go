package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("traceId")
		if traceId == "" {
			traceId = uuid.New().String()
		}
		c.Set("traceId", traceId)
		log.Println("traceId generated: ", traceId)
		c.Next()
	}
}
