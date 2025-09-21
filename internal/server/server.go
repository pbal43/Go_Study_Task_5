package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"toDoList/internal"
)

type UserStorage interface {
	GetAllUsers()
	SaveUser(user userDomain.User) error
	GetUser(userReq userDomain.UserRequest) (userDomain.User, error)
	UpdateUser()
}

type TaskStorage interface {
	GetAllTasks()
	GetOneTaskByID()
	AddTask()
	UpdateTask()
}

type Storage interface {
	UserStorage
	TaskStorage
}

type ToDoListApi struct {
	srv         *http.Server
	db          Storage
	tokenSigner auth.HS256Signer
}

func NewServer(cfg internal.Config, db Storage) *ToDoListApi {

	signer := auth.HS256Signer{
		Secret:     []byte("ultraSecretKey123"),
		Issuer:     "todolistService",
		Audience:   "todolistClient",
		AccessTTL:  15 * time.Minute,
		RefreshTTL: 24 * 7 * time.Hour,
	}

	HttpSrv := http.Server{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	api := ToDoListApi{srv: &HttpSrv, db: db, tokenSigner: signer}

	api.configRouter()

	return &api
}

func (api *ToDoListApi) Run() error {
	return api.srv.ListenAndServe()
}

func (api *ToDoListApi) ShutDown() error {
	return nil
}

func (api *ToDoListApi) configRouter() {
	router := gin.Default()

	tasks := router.Group("/tasks")
	{
		tasks.GET("/", GetTasks)
		tasks.GET("/:id", GetTaskByID)
		tasks.POST("/", CreateTask)
		tasks.PUT("/:id", UpdateTask)
		tasks.DELETE("/:id", DeleteTask)
	}

	users := router.Group("/users")
	{
		users.GET("/", GetAllUsers)
		users.GET("/:id", GetUserByID)
		users.POST("/register", Register)
		users.POST("/login", Login)
		users.PUT("/:id", UpdateUser)
		users.DELETE("/:id", DeleteUser)
	}

	api.srv.Handler = router
}
