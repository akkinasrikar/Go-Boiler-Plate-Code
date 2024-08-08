package middleware

import (
	"context"

	"example/boiler-plate/constants"
	"example/boiler-plate/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateJwtAuthToken() func(*gin.Context) {
	return func(ctx *gin.Context) {
		var header models.AuthData
		ctx.ShouldBindHeader(&header)
		Ctx := SetRequestContext(header)

		if header.Authorization == "" {
			ecomErr := setAuthError(constants.ErrorMessage.AUTHENTICATION_REQUIRED)
			ctx.JSON(constants.ErrorType.AUTHENTICATION_ERROR.Code, ecomErr)
			ctx.Abort()
			return
		}
		secret_key := "testing"
		token, err := jwt.Parse(header.Authorization, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret_key), nil
		})
		if err != nil {
			ecomErr := setAuthError(constants.ErrorMessage.INVALID_AUTHENTICATION)
			ctx.JSON(constants.ErrorType.AUTHENTICATION_ERROR.Code, ecomErr)
			ctx.Abort()
			return
		}
		if !token.Valid {
			ecomErr := setAuthError(constants.ErrorMessage.AUTHENTICATION_REQUIRED)
			ctx.JSON(constants.ErrorType.AUTHENTICATION_ERROR.Code, ecomErr)
			ctx.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		Ctx = setEcomCtx(ctx, claims, header)
		ctx.Set("AppCtx", Ctx)
		ctx.Next()
	}
}

func SetRequestContext(authData models.AuthData) context.Context {
	var ecomCtx context.Context
	ecomCtx = context.Background()
	ecomCtx = context.WithValue(ecomCtx, models.AppCtxKey("AuthData"), authData)
	return ecomCtx
}

func setAuthError(errMsg string) models.AppErrorResponse {
	zwErrBody := models.AppErrorBody{
		Message: errMsg,
		Type:    constants.ErrorType.AUTHENTICATION_ERROR.Name,
	}
	return models.AppErrorResponse{
		ErrorType: zwErrBody,
	}
}

func setEcomCtx(ecomCtx context.Context, tokenClaims jwt.Claims, header models.AuthData) context.Context {
	var authData models.AuthData

	authData.Authorization = header.Authorization
	authData.ISsandBox = header.ISsandBox
	sub := tokenClaims.(jwt.MapClaims)["sub"].(string)
	usersId := tokenClaims.(jwt.MapClaims)["usersId"].(float64)
	authData.UserName = sub
	authData.UsersId = int64(usersId)
	ecomCtxUpdated := context.Background()
	ecomCtxUpdated = context.WithValue(ecomCtxUpdated, models.AppCtxKey("AuthData"), authData)
	return ecomCtxUpdated
}
