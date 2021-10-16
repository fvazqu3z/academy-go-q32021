package app

import (
	"fmt"
	"log"
	"wizeline/pool"

	"wizeline/common"
	"wizeline/controller"
	"wizeline/repository"
	"wizeline/routes"
	"wizeline/service"
	"wizeline/usecase"
)

//Start run the server
func Start() {
	var connectionPoll *pool.GoroutinePool = pool.NewGoroutinePool(common.PoolSize)
	var userRepository repository.UserRepository = repository.NewRepository()
	var userService service.Service = service.NewUserService(userRepository)
	var userUseCase usecase.UseCase = usecase.NewUseCase(userService)
	var userController,connectionPool = controller.NewController(userUseCase, connectionPoll)
	var userRouter routes.Router = routes.NewMuxRouter()

	userRouter.GET("/", userController.HomeController)
	userRouter.GET("/status", userController.StatusController)
	userRouter.GET("/users", userController.GetUsers)
	userRouter.GET("/saveusers", userController.SaveUsers)

	err := userRouter.START(common.Port)

	if err == nil {
		fmt.Printf("Server was running on port %v \n", common.Port)
		log.Println("Server was running on port " +  common.Port)
	}

	if err != nil {
		log.Fatal("Error initializing server")
	}
	connectionPool.Close()
}
