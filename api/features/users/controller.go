package User

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Call userService to userController
type UserController struct {
	userService UserService
}

// Create new Instance from UserController
func NewUserController(userService UserService) *UserController {
	return &UserController{userService: userService}
}

// calling api
func (c *RegisterUser) RegisterUser(ctx *gin.Context) {
	var userInput RegisterUser
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.RegisterUser(&userInput)
	if err != nil {
		// Mengembalikan kesalahan jika registrasi gagal
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan registrasi pengguna"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Registrasi pengguna berhasil"})
}
