package repositories

import (
	"example/boiler-plate/database"
	"example/boiler-plate/models"
	"example/boiler-plate/models/entities"
)

type Repository struct {
	dbStore database.DB
}

func NewRepository(dbStore database.DB) RepositoryInterface {
	return &Repository{
		dbStore: dbStore,
	}
}

type RepositoryInterface interface {
	SignUp(userDetails entities.SignUp) (entities.SignUp, models.AppError)
	Login(userDetails entities.Login) (entities.SignUp, models.AppError)
}
