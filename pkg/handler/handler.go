package handler

import (
	"github.com/ellywynn/rest-school/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	root := router.Group("/")
	{
		root.GET("/", h.indexPage)
	}

	return router
}
