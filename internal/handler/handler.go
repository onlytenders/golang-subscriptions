package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/onlytenders/golang-subscriptions/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	services service.SubscriptionService
	logger   *zap.Logger
}

func NewHandler(services service.SubscriptionService, logger *zap.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		subscriptions := api.Group("/subscriptions")
		{
			subscriptions.POST("/", h.createSubscription)
			subscriptions.GET("/:id", h.getSubscriptionByID)
			subscriptions.PUT("/:id", h.updateSubscription)
			subscriptions.DELETE("/:id", h.deleteSubscription)
			subscriptions.GET("/", h.listSubscriptions)
			subscriptions.GET("/total", h.getTotalCost)
		}
	}

	return router
}
