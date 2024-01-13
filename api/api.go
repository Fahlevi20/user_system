package api

import (
	User "github.com/Fahlevi20/user_system/api/features/users"
	"github.com/gin-gonic/gin"
)

type Router interface {
	SetupRouter(router *gin.Engine)
}

type router struct{}

func NewRouter() Router {
	return &router{}
}

func (r *router) SetupRouter(router *gin.Engine) {
	// Inisialisasi database
	db := initializeDatabase()

	// Inisialisasi repository dan service
	userRepository := User.NewUserRepository(db)
	userService := User.NewUserService(userRepository)

	// Inisialisasi controller
	userController := User.NewUserController(userService)

	// Grup router untuk endpoint API
	api := router.Group("/api")

	// Rute untuk registrasi pengguna
	api.POST("/register", userController.RegisterUserAPI)
}
