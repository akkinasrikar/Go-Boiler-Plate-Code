package controllers

import (
	"context"
	"net/http"

	"example/boiler-plate/models"
	services "example/boiler-plate/services/login"
	"example/boiler-plate/validators/helper"
	validator "example/boiler-plate/validators/login"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginService   services.LoginService
	loginValidator validator.LoginValidator
}

func NewLoginHandler(loginService services.LoginService, loginValidator validator.LoginValidator) *LoginHandler {
	return &LoginHandler{
		loginService:   loginService,
		loginValidator: loginValidator,
	}
}

func (lh *LoginHandler) SignUp(ctx *gin.Context) {
	req, err := lh.loginValidator.ValidateSignUp(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	resp, err := lh.loginService.SignUp(req)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (lh *LoginHandler) Login(ctx *gin.Context) {
	req, err := lh.loginValidator.ValidateLogin(ctx)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	resp, err := lh.loginService.Login(req)
	if err.Message != nil {
		err := helper.SetInternalError(err.Message.Error())
		ctx.JSON(int(err.ErrorType.Code), &err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (lh *LoginHandler) HomePage(ctx *gin.Context) {
	ecomGinCtx, _ := ctx.Get("AppCtx")
	appCtx := ecomGinCtx.(context.Context)
	authData := appCtx.Value(models.AppCtxKey("AuthData")).(models.AuthData)
	ctx.JSON(200, gin.H{
		"message": "Welcome " + authData.UserName + " to HomePage",
	})
}
