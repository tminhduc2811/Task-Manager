package task_manager
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
	dbConnect, err := mgo.Dial(config.DatabaseName)
	if err != nil {
		println(err)
		panic(err)
	}
	taskRepo := repository.NewMgoTaskRepository(dbConnect, config.DatabaseName, config.TaskCollection)

	router := gin.Default()
	tasks := router.Group("/tasks")
	{
		routers.TaskRoutes(tasks, *controllers.NewTaskService(taskRepo))
	}
	fmt.Println("Listening at 8080")
	router.Run(":8080")
}