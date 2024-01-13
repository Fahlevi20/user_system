package api

import "github.com/gin-gonic/gin"

type Router interface {
	SetupRouter(router *gin.Engine)
}

type router struct{}

func NewRouter() Router {
	return &router{}
}

func (r *router) SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	db = initializeDatabase()

	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)

	userController := controller.NewUserController(userService)

	router := gin.Default()

	router.POST("/register", RegisterUser)
	return router

}
