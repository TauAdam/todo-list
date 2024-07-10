package handler

import (
	"github.com/TauAdam/todo-list/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.sighUp)
		auth.POST("/sign-in", h.sighIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getLists)
			lists.POST("/", h.createList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateListById)
			lists.DELETE("/:id", h.deleteListById)

			items := lists.Group(":id/items")
			{
				items.GET("/", h.getItems)
				items.POST("/", h.createItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItemById)
				items.DELETE("/:item_id", h.deleteItemById)
			}
		}
	}
	return router
}
