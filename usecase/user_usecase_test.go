package usecase

import (
	"log"
	"os"
	"testing"

	"wizeline/common"
	"wizeline/repository"
	"wizeline/service"

	"github.com/go-resty/resty/v2"
)
func TestUserUseCase_GetUsers(t *testing.T) {
	common.CSVFileNameInput  = "./../data/users.csv"
	common.CSVFileNameOutput  = "./../data/localusers.csv"

	var userRepository repository.UserRepository = repository.NewRepository()
	var userService service.Service = service.NewUserService(userRepository)
	var userUseCase UseCase = NewUseCase(userService)

	userUseCase.GetUsers()
}

func TestUserUseCase_SaveUsers(t *testing.T) {
	common.CSVFileNameInput  = "./../data/users.csv"
	common.CSVFileNameOutput  = "./../data/localusers.csv"

	var userRepository repository.UserRepository = repository.NewRepository()
	var userService service.Service = service.NewUserService(userRepository)
	var userUseCase UseCase = NewUseCase(userService)

	client := resty.New()
	userUseCase.SaveUsers(client)

	localData, localError := os.ReadFile(common.CSVFileNameInput)
	remoteData, remoteError := os.ReadFile(common.CSVFileNameOutput)

	if localError != nil{
		log.Println(localError)
		t.Errorf("Error reading local users")
	}

	if remoteError != nil{
		log.Println(remoteError)
		t.Errorf("Error reading remote users")
	}

	if string(localData) != string(remoteData){
		t.Errorf("The local data do not match with server data")
	}


}
