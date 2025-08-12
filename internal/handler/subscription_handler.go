package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onlytenders/golang-subscriptions/internal/models"
)

type createSubscriptionInput struct {
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	ServiceName string    `json:"service_name" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	StartDate   string    `json:"start_date" binding:"required"`
	EndDate     *string   `json:"end_date"`
}

// @Summary Create a subscription
// @Description Create a new subscription
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Param   input     body    createSubscriptionInput  true  "Subscription Info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions [post]
func (h *Handler) createSubscription(c *gin.Context) {
	var input createSubscriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("01-2006", input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, expected MM-YYYY"})
		return
	}

	var endDate *time.Time
	if input.EndDate != nil {
		parsedEndDate, err := time.Parse("01-2006", *input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, expected MM-YYYY"})
			return
		}
		endDate = &parsedEndDate
	}

	sub := &models.Subscription{
		UserID:      input.UserID,
		ServiceName: input.ServiceName,
		Price:       input.Price,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	id, err := h.services.Create(c.Request.Context(), sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// @Summary Get a subscription by ID
// @Description Get a subscription by its ID
// @Tags subscriptions
// @Produce  json
// @Param   id     path    string  true  "Subscription ID"
// @Success 200 {object} models.Subscription
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [get]
func (h *Handler) getSubscriptionByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	sub, err := h.services.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sub)
}

// @Summary Update a subscription
// @Description Update an existing subscription
// @Tags subscriptions
// @Accept  json
// @Produce  json
// @Param   id     path    string  true  "Subscription ID"
// @Param   input     body    createSubscriptionInput  true  "Subscription Info"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [put]
func (h *Handler) updateSubscription(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input createSubscriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("01-2006", input.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, expected MM-YYYY"})
		return
	}

	var endDate *time.Time
	if input.EndDate != nil {
		parsedEndDate, err := time.Parse("01-2006", *input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, expected MM-YYYY"})
			return
		}
		endDate = &parsedEndDate
	}

	sub := &models.Subscription{
		ID:          id,
		UserID:      input.UserID,
		ServiceName: input.ServiceName,
		Price:       input.Price,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := h.services.Update(c.Request.Context(), sub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Delete a subscription
// @Description Delete a subscription by its ID
// @Tags subscriptions
// @Param   id     path    string  true  "Subscription ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [delete]
func (h *Handler) deleteSubscription(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.services.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// @Summary List all subscriptions
// @Description Get a list of all subscriptions
// @Tags subscriptions
// @Produce  json
// @Success 200 {array} models.Subscription
// @Failure 500 {object} map[string]string
// @Router /subscriptions [get]
func (h *Handler) listSubscriptions(c *gin.Context) {
	subs, err := h.services.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subs)
}

// @Summary Get total cost of subscriptions
// @Description Get the total cost of subscriptions for a user in a given month and year
// @Tags subscriptions
// @Produce  json
// @Param   user_id     query    string  true  "User ID"
// @Param   service_name query    string  false  "Service Name"
// @Param   year     query    int  true  "Year"
// @Param   month     query    int  true  "Month"
// @Success 200 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/total [get]
func (h *Handler) getTotalCost(c *gin.Context) {
	userID, err := uuid.Parse(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	serviceName := c.Query("service_name")

	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year"})
		return
	}

	month, err := strconv.Atoi(c.Query("month"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month"})
		return
	}

	total, err := h.services.TotalCost(c.Request.Context(), userID, serviceName, year, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_cost": total})
}
