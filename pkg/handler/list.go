package handler

import (
	todolist "github.com/TauAdam/todo-list"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary		Create a new list
// @Tags			lists
// @Description	create a new list
// @Security		ApiKeyAuth
// @ID				create-list
// @Accept			json
// @Produce		json
// @Param			input	body		todolist.TodoList	true	"List info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := provideUserId(c)
	if err != nil {
		return
	}
	var input todolist.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getListsResponse struct {
	Data []todolist.TodoList `json:"data"`
}

// @Summary		Get All Lists
// @Security		ApiKeyAuth
// @Tags			lists
// @Description	get all lists
// @ID				get-all-lists
// @Accept			json
// @Produce		json
// @Success		200		{object}	getListsResponse
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/api/lists [get]
func (h *Handler) getLists(c *gin.Context) {
	userId, err := provideUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := provideUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateListById(c *gin.Context) {
	userId, err := provideUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input todolist.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.TodoList.Update(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{"ok"})
}

// @Summary		Delete list
// @Security		ApiKeyAuth
// @Tags			lists
// @Description	Delete list by id
// @ID				delete-list
// @Produce		json
// @Param			id		path		int	true	"List ID"
// @Success		200		{object}	StatusResponse
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Router			/api/lists/{id} [delete]
func (h *Handler) deleteListById(c *gin.Context) {
	userId, err := provideUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}
	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
