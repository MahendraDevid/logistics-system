package functional

import (
	"context"
	"testing"

	"pricing-service/internal/domain"
	"pricing-service/internal/service"
	"pricing-service/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePricing(t *testing.T) {

	repo := mocks.NewMockPricingRepository(nil)

	svc := service.NewPricingService(repo)

	req := domain.CalculationRequest{
		Origin:      "Jakarta",
		Destination: "Bandung",

		WeightKG: 5,

		Length: 20,
		Width:  20,
		Height: 20,

		ServiceType: "REGULAR",
	}

	result, err := svc.CalculateTariff(
		context.Background(),
		req,
	)

	assert.NoError(t, err)

	assert.NotNil(t, result)

	assert.Greater(t, result.Total, 0.0)
}