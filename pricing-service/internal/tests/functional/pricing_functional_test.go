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
	OriginPostalCode:      "10110",
	DestinationPostalCode: "40115",

	WeightKG: 2,

	LengthCM: 10,
	WidthCM:  10,
	HeightCM: 10,

	ServiceType:  "REG",
	UseInsurance: true,
}

	resp := svc.CalculateTariff(req)

	assert.NotNil(t, resp)
	assert.Equal(t, "2-3 Days", resp.Estimated)
}