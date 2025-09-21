package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func GetAllUsers(ctx *gin.Context) {}

func GetUserByID(ctx *gin.Context) {}

func (srv *ToDoListApi) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usecase := userusecase.NewUserUsecase(srv.db)
	if err := usecase.SaveUser(user); err != nil {
		if errors.Is(err, userErrors.ErrorUserIsAlreadyExist) {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(ctx *gin.Context) {}

func UpdateUser(ctx *gin.Context) {}

func DeleteUser(ctx *gin.Context) {}
