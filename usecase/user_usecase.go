package usecase

import (
	"wizeline/model"
	"wizeline/service"

	"github.com/go-resty/resty/v2"
)

//UseCase interface to manage use cases
type UseCase interface{
	GetUsers() (string, error)
	GetUser(userID int) (model.User, error)
	SaveUsers(client *resty.Client) error
}

//UserUseCase struct to manage use cases
type UserUseCase struct {
	useCase UseCase
}

var userService service.Service

//NewUseCase method to create a new use case
func NewUseCase(interfaceService service.Service) UseCase {
	userService = interfaceService
	return &UserUseCase{}
}

//GetUsers retrieve all users
func (userUseCase *UserUseCase) GetUsers() (string, error) {
	return  userService.GetUsers()
}

//GetUser retrieve a specific user
func (userUseCase *UserUseCase) GetUser(userID int) (model.User, error) {
	return userService.GetUser(userID)
}

//SaveUsers save users to csv file
func (userUseCase *UserUseCase) SaveUsers(client *resty.Client)  error {
	return userService.SaveUsers(client)
}
