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

func NewPricingHandler(s *service.PricingService) *PricingHandler {
	return &PricingHandler{
		service: s,
	}
}

func (h *PricingHandler) Calculate(c *gin.Context) {

	var req domain.PricingRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	result, err := h.service.Calculate(req)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, result)
}