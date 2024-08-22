package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

type NameHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "First end point is now up, wow!!!")
}

func NewNameHandler() *NameHandler {
	return &NameHandler{}
}

func (n *NameHandler) Name(c *gin.Context) {
	name := c.Params.ByName("id")
	if name != "" {
		c.JSON(http.StatusOK, fmt.Sprintf("Your name is %s", name))
	}
	c.JSON(404, "Can not find your name in request params")
}
