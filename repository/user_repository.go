package repository

import "wizeline/model"

//UserRepository interface to manage user repository
type UserRepository interface{
	GetUsers() (string, error)
	GetUser(userID int) (model.User, error)
	SaveUsersToDataSource(users []model.User) error
	GetUsersFromDataSource(data [][]string) ([]model.User, error)
}
