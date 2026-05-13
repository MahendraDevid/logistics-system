package service_test

import (
	"testing"

	"pricing-service/internal/domain"
	"pricing-service/internal/mocks"
	"pricing-service/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTariff(t *testing.T) {

	mockRepo := mocks.NewMockPricingRepository()

	pricingSvc := service.NewPricingService(mockRepo)

	req := domain.CalculationRequest{
		Origin:      "Jakarta",
		Destination: "Bandung",
		Weight:      2,
		Length:      20,
		Width:       20,
		Height:      20,
		ServiceType: "REG",
	}

	resp := pricingSvc.CalculateTariff(req)

	assert.Equal(t, 24000.0, resp.Total)
	assert.Equal(t, "2-3 Days", resp.Estimated)
}