package server

import (
	"fmt"
	"net/http"
	"time"
	"toDoList/internal"
	"toDoList/internal/domain/user/user_models"
	auth "toDoList/internal/server/auth/user_auth"
	"toDoList/internal/server/middleware"

	"github.com/gin-gonic/gin"
)

type UserStorage interface {
	GetAllUsers() []user_models.User
	SaveUser(user user_models.User) (user_models.User, error)
	GetUserByID(userID string) (user_models.User, error)
	GetUserByEmail(email string) (user_models.User, error)
	UpdateUser(user user_models.User) (user_models.User, error)
	DeleteUser(userID string) error
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
		tasks.GET("/", api.getTasks)
		tasks.GET("/:id", api.getTaskByID)
		tasks.POST("/", api.createTask)
		tasks.PUT("/:id", api.updateTask)
		tasks.DELETE("/:id", api.deleteTask)
	}

	users := router.Group("/users")
	{
		users.GET("/", api.getAllUsers)
		users.GET("/:id", api.getUserByID)
		users.POST("/register", api.register)
		users.POST("/login", api.login)
		users.POST("/admin-login", api.loginAdmin)
		users.PUT("/:id", middleware.AuthMiddleware(api.tokenSigner), api.updateUser)
		users.DELETE("/:id", middleware.AuthMiddleware(api.tokenSigner), api.deleteUser)
	}

	api.srv.Handler = router
}
