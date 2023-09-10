package repository

import (
	"app-aposrgb/src/database"
	"app-aposrgb/src/models"
)

func GetUsers() []models.User {
	users := []models.User{}
	database.DB.Find(&users)
	return users
}

func GetUser(id int) (models.User, bool) {
	user := models.User{}
	database.DB.Model(models.User{Id: id}).First(&user)
	if user.Id != id {
		return user, false
	}
	return user, true
}
