package main
import (
	"./common"
	"./controllers"
	"./repository"
	"./routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main()  {
	config, err := common.ReadConfig("./config.yml")
	if err != nil {
		panic(err)
	}
	dbConnect, err := mgo.Dial(config.DatabaseAddr)
	if err != nil {
		println(err)
		panic(err)
	}
	taskRepo := repository.NewMgoTaskRepository(dbConnect, config.DatabaseName, config.TaskCollection)
	userRepo := repository.NewMgoUserRepository(dbConnect, config.DatabaseName, config.UserCollection)

	router := gin.Default()
	tasks := router.Group("/tasks")
	{
		routers.TaskRoutes(tasks, *controllers.NewTaskController(taskRepo))
	}
	users := router.Group("/users")
	{
		routers.UserRoutes(users, *controllers.NewUserController(userRepo))
	}

	fmt.Println("Listening at 8080")
	router.Run(":8080")
}