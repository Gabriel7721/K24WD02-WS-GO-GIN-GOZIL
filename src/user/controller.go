package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Repo *Repository
}

func NewController(repo *Repository) *Controller {
	return &Controller{Repo: repo}
}

func (ctrl Controller) Register(ctx *gin.Context) {
	var input User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}

	ctrl.Repo.Create(&input)
	ctx.JSON(http.StatusOK, gin.H{"message": "Register success!"})
}
