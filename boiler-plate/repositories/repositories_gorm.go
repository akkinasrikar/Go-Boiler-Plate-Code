package repositories

import (
	"example/boiler-plate/models"
	"example/boiler-plate/models/entities"
	"example/boiler-plate/validators/helper"
	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) SignUp(userDetails entities.SignUp) (entities.SignUp, models.AppError) {
	if err := r.dbStore.Create(&userDetails); err != nil {
		return entities.SignUp{}, *helper.ErrorInternalSystemError("Error while signing up : " + err.Error())
	}
	return userDetails, models.AppError{}
}

func (r *Repository) Login(userDetails entities.Login) (entities.SignUp, models.AppError) {
	var user entities.SignUp
	_, err := r.dbStore.Where("user_name = ? OR user_email = ?", userDetails.Name, userDetails.Name).Find(&user)
	if err != nil {
		return entities.SignUp{}, *helper.ErrorInternalSystemError(err.Error())
	}
	if user.Name == "" {
		return entities.SignUp{}, *helper.ErrorInternalSystemError("User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password))
	if err != nil {
		return entities.SignUp{}, *helper.ErrorInternalSystemError("Password is incorrect")
	}
	return user, models.AppError{}
}
