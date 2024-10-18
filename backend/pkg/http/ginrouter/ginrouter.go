package ginrouter

import "github.com/gin-gonic/gin"

type Router struct {
	R *gin.Engine
}

// New create new Gin router
func New() *Router {
	router := gin.Default()
	return &Router{R: router}
}
