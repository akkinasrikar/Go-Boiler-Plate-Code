package login

import (
	"example/boiler-plate/models"
	"github.com/gin-gonic/gin"
)

type loginValidator struct{}

func NewValidator() LoginValidator {
	return &loginValidator{}
}

type LoginValidator interface {
	ValidateSignUp(ctx *gin.Context) (models.SignUp, models.AppError)
	ValidateLogin(ctx *gin.Context) (models.Login, models.AppError)
}
