package main

import (
	"github.com/gin-gonic/gin"

	"pricing-service/internal/handler"
	"pricing-service/internal/service"
)

func main() {

	r := gin.Default()

	pricingService := service.NewPricingService()

	pricingHandler := handler.NewPricingHandler(pricingService)

	r.POST("/calculate", pricingHandler.Calculate)

	r.Run(":8080")
}