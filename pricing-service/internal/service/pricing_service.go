package service

import "pricing-service/internal/domain"

type PricingRepository interface {
	GetBaseRate(origin, destination, serviceType string) float64
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
	req domain.CalculationRequest,
) domain.CalculationResponse {

	actualWeight := req.Weight
	volumetric := (req.Length * req.Width * req.Height) / 6000

	finalWeight := actualWeight

	if volumetric > actualWeight {
		finalWeight = volumetric
	}

	baseRate := s.repo.GetBaseRate(
		req.Origin,
		req.Destination,
		req.ServiceType,
	)

	base := finalWeight * baseRate
	insurance := base * 0.02

	total := base + insurance

	return domain.CalculationResponse{
		BaseTariff: base,
		Insurance:  insurance,
		Discount:   0,
		Total:      total,
		Estimated:  "2-3 Days",
	}
}