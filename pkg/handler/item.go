package handler

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := provideUserId(c)
	if err != nil {
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input todolist.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": itemId,
	})
}
func (h *Handler) getItems(c *gin.Context) {

}

func (h *Handler) getItemById(c *gin.Context) {

}

func (h *Handler) updateItemById(c *gin.Context) {

}

func (h *Handler) deleteItemById(c *gin.Context) {

}
