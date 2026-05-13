package service

import (
	"context"
	"pricing-service/internal/domain"
)

type PricingRepository interface {
	GetBaseRate(serviceType string) float64
}

type PricingService struct {
	repo PricingRepository
}

func NewPricingService(repo PricingRepository) *PricingService {
	return &PricingService{
		repo: repo,
	}
}

func (s *PricingService) CalculateTariff(
	ctx context.Context,
	req domain.CalculationRequest,
) (domain.CalculationResponse, error) {

	actualWeight := req.WeightKG

	volumetric :=
		(req.Length * req.Width * req.Height) / 6000

	finalWeight := actualWeight

	if volumetric > actualWeight {
		finalWeight = volumetric
	}

	rate := s.repo.GetBaseRate(req.ServiceType)

	if rate == 0 {
		rate = 10000
	}

	base := finalWeight * rate

	insurance := base * 0.02

	total := base + insurance

	return domain.CalculationResponse{
		BaseTariff: base,
		Insurance:  insurance,
		Discount:   0,
		Total:      total,
		Estimated:  "2-3 Days",
	}, nil
}