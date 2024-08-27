package middlewares

import (
	"fmt"
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/soheilsirousi/golang-web-api/src/api/helper"
)

func Limiter() gin.HandlerFunc {
	lim := tollbooth.NewLimiter(1, nil)

	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lim, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.NewBaseResponseWithError("", false, -100,
				fmt.Errorf("too many request")))
			return
		}
		c.Next()
	}
}
