package handler

import "github.com/gin-gonic/gin"

func (h *Handler) createList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateListById(c *gin.Context) {

}

func (h *Handler) deleteListById(c *gin.Context) {

}
