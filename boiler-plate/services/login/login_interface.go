package services

import (
	"example/boiler-plate/models"
	"example/boiler-plate/models/responses"
	"example/boiler-plate/repositories"
	"github.com/go-redis/redis"
)

type loginService struct {
	repoService repositories.RepositoryInterface
	redisClient *redis.Client
}

func NewLoginService(respoService repositories.RepositoryInterface, redisClient *redis.Client) LoginService {
	return &loginService{
		repoService: respoService,
		redisClient: redisClient,
	}
}

type LoginService interface {
	SignUp(req models.SignUp) (responses.SingUp, models.AppError)
	Login(req models.Login) (responses.Login, models.AppError)
}
