package handler

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) sighUp(c *gin.Context) {
	var input todolist.User
	if err := c.BindJSON(&input); err != nil {
		NewError(c, http.StatusBadRequest, err.Error())
		return
	}
}
func (h *Handler) sighIn(c *gin.Context) {

}
