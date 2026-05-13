package functional

import (
	"testing"

	"pricing-service/internal/domain"
	"pricing-service/internal/mocks"
	"pricing-service/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePricing(t *testing.T) {

	repo := mocks.NewMockPricingRepository()

	svc := service.NewPricingService(repo)

	req := domain.CalculationRequest{
		Origin:      "Jakarta",
		Destination: "Bandung",
		Weight:      1,
		Length:      10,
		Width:       10,
		Height:      10,
		ServiceType: "REG",
	}

	resp := svc.CalculateTariff(req)

	assert.NotNil(t, resp)
	assert.Equal(t, "2-3 Days", resp.Estimated)
}