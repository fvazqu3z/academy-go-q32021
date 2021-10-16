package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"wizeline/common"
	"wizeline/model"
	"wizeline/repository"

	"github.com/go-resty/resty/v2"
)

//Service interface to manage services
type Service interface {
	GetUsers() (string, error)
	GetUser(userID int) (model.User, error)
	SaveUsers(client *resty.Client) error
}

//UserService struct to manage services
type UserService struct {
	service Service
}

var repo repository.UserRepository

//NewUserService method to create a new user service
func NewUserService(interfaceUserRepository repository.UserRepository) Service {
	repo = interfaceUserRepository
	return &UserService{}
}

//GetUsers retrieve all users
func (interfaceService *UserService) GetUsers() (string, error){
	return repo.GetUsers()
}

//GetUser retrieve a specific user
func (interfaceService *UserService) GetUser(userID int) (model.User, error){
	return repo.GetUser(userID)
}

// SaveUsers Get all user from server API and store it in local file
func (interfaceService *UserService) SaveUsers(client *resty.Client) error {
	resp, err := client.R().Get(common.ServerUsersEndPoint)

	if err != nil {
		log.Println(err)
		return errors.New("error getting users")
	}

	// Unmarshal JSON data
	var jsonData []model.User
	textBytes := []byte(resp.Body())
	err = json.Unmarshal(textBytes, &jsonData)

	if err != nil {
		log.Println(err)
		return errors.New("error unmarshalling users")
	}

	err = repo.SaveUsersToDataSource(jsonData)

	if err != nil {
		log.Println(err)
		return errors.New("error saving users")
	}

	fmt.Println("Data was successfully saved")
	return nil

}
