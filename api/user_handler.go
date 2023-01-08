package api

import (
	"net/http"

	"github.com/fajarhadifirmansyah/bedbo/store"
	"github.com/fajarhadifirmansyah/bedbo/types"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	store store.UserStorer
}

func NewUserHandler(uStore store.UserStorer) *UserHandler {
	return &UserHandler{
		store: uStore,
	}
}

func (uc *UserHandler) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*types.User)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": currentUser}})
}
