package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pricing-service/internal/domain"
	"pricing-service/internal/service"
)

type PricingHandler struct {
	service *service.PricingService
}

func NewPricingHandler(
	service *service.PricingService,
) *PricingHandler {

	return &PricingHandler{
		service: service,
	}
}

func (h *PricingHandler) CalculatePricing(
	c *gin.Context,
) {

	var req domain.CalculationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := h.service.CalculateTariff(
		c.Request.Context(),
		req,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}