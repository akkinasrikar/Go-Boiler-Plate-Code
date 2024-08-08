package server

import (
	"example/boiler-plate/controllers"
	"example/boiler-plate/database"
	"example/boiler-plate/middleware"
	"example/boiler-plate/repositories"
	servicesLogin "example/boiler-plate/services/login"
	validatorsLogin "example/boiler-plate/validators/login"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func setUpRoutes(router *gin.Engine, db database.DB, redisClient *redis.Client) {
	ecomStore := repositories.NewRepository(db)
	validatorsLogin := validatorsLogin.NewValidator()

	servicesLogin := servicesLogin.NewLoginService(ecomStore, redisClient)
	LoginHandler := controllers.NewLoginHandler(servicesLogin, validatorsLogin)
	loginHandler(router, *LoginHandler)

}

func loginHandler(router *gin.Engine, LoginHandler controllers.LoginHandler) {
	router.POST("/signup", LoginHandler.SignUp)
	router.POST("/login", LoginHandler.Login)

	router.Use(middleware.ValidateJwtAuthToken())
	router.Use((middleware.TraceIDMiddleware()))
	router.GET("/homePage", LoginHandler.HomePage)
}
