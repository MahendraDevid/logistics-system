package functional

import (
	"testing"

	"pricing-service/internal/domain"
	"pricing-service/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePricing(t *testing.T) {

	svc := service.NewPricingService()

	req := domain.PricingRequest{
		WeightKG:     2,
		LengthCM:     20,
		WidthCM:      20,
		HeightCM:     20,
		ServiceType:  "REGULER",
		UseInsurance: true,
		PromoCode:    "HEMAT10",
	}

	result, err := svc.Calculate(req)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.TotalPayment > 0)
}