package controller

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"wizeline/common"

	"wizeline/errors"
	"wizeline/pool"
	"wizeline/usecase"

	"github.com/go-resty/resty/v2"
)

// UserController Struct to manage user controller
type UserController struct{

}

// Controller interface to give method access
type Controller interface {
	GetUsers(response http.ResponseWriter, request *http.Request)
	GetUser(response http.ResponseWriter, request *http.Request)
	HomeController(response http.ResponseWriter, request *http.Request)
	StatusController(response http.ResponseWriter, request *http.Request)
	SaveUsers(response http.ResponseWriter, request *http.Request)
}

type requestTask struct {
	ID int
	Task func(...interface{})
}

func (t *requestTask) Run(){
	t.Task(t.ID)
}


var userUseCase usecase.UseCase
var connectionPoll *pool.GoroutinePool
var wg  *sync.WaitGroup
var currentTask int

//NewController create controller
func NewController(interfaceUserUseCase usecase.UseCase, executorPool *pool.GoroutinePool) (Controller,*pool.GoroutinePool) {
	userUseCase = interfaceUserUseCase
	connectionPoll = executorPool
	wg = &sync.WaitGroup{}
	currentTask = 0
	return &UserController{} , connectionPoll
}


//GetUsers retrieve all users
func (interfaceController *UserController) GetUsers(response http.ResponseWriter, request *http.Request) {
	wg.Add(common.Delta)
	currentTask++

	taskToSchedule := func(p ...interface{}) {
		response.WriteHeader(http.StatusOK)
		response.Header().Set("Content-Type", "application/json")
		users, err := userUseCase.GetUsers()

		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			responseError := errors.NewResponseError(http.StatusInternalServerError, "Error reading users csv file", err)
			http.Error(response, responseError.Error(), http.StatusInternalServerError)
			wg.Done()
			currentTask--
			return
		}
		_, err = fmt.Fprintf(response, users)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			responseError := errors.NewResponseError(http.StatusInternalServerError, "Error getting user", err)
			http.Error(response, responseError.Error(), http.StatusInternalServerError)
			wg.Done()
			currentTask--
			return
		}
		wg.Done()
		currentTask--
	}

	t := &requestTask{
		ID:   currentTask,
		Task: taskToSchedule,
	}
	connectionPoll.ScheduleWorks(t)
	wg.Wait()
	fmt.Println("Request #Get Users# finished")

}

//SaveUsers save the users in csv file
func (interfaceController *UserController) SaveUsers(response http.ResponseWriter, request *http.Request) {
	client := resty.New()
	err := userUseCase.SaveUsers(client)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		responseError := errors.NewResponseError(http.StatusInternalServerError,"Error consulting users API" , err)
		http.Error(response,responseError.Error(),http.StatusInternalServerError)
		return
	}
}


func (interfaceController *UserController) GetUser(response http.ResponseWriter, request *http.Request) {}

//HomeController display Home message
func  (interfaceController *UserController)  HomeController(response http.ResponseWriter, request *http.Request) {
	wg.Add(common.Delta)
	currentTask++

	taskToSchedule := func(p ...interface{}) {
		response.WriteHeader(http.StatusOK)
		_, err := fmt.Fprintf(response, "Home")
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			responseError := errors.NewResponseError(http.StatusInternalServerError, "Error displaying home", err)
			http.Error(response, responseError.Error(), http.StatusInternalServerError)
			wg.Done()
			currentTask--
			return
		}
		wg.Done()
		currentTask--
	}

	t := &requestTask{
		ID:   currentTask,
		Task: taskToSchedule,
	}
	connectionPoll.ScheduleWorks(t)
	wg.Wait()
	fmt.Println("Request #Get User# finished")

}

//StatusController return running = true when all is ok , or error is other cases
func  (interfaceController *UserController)  StatusController(response http.ResponseWriter, request *http.Request) {
	wg.Add(common.Delta)
	currentTask++

	taskToSchedule := func(p ...interface{}) {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		_, err := io.WriteString(response, `{"Running": true}`)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			responseError := errors.NewResponseError(http.StatusInternalServerError, "Server is not running", err)
			http.Error(response, responseError.Error(), http.StatusInternalServerError)
			wg.Done()
			currentTask--
			return
		}
		wg.Done()
		currentTask--
	}

	t := &requestTask{
		ID:   currentTask,
		Task: taskToSchedule,
	}
	connectionPoll.ScheduleWorks(t)
	wg.Wait()
	fmt.Println("Request #Status Controller# finished")


}