package service

import (
	"errors"
	"math"
	"strings"

	"pricing-service/internal/domain"
)

type PricingService struct {
}

func NewPricingService() *PricingService {
	return &PricingService{}
}

func (s *PricingService) Calculate(req domain.PricingRequest) (*domain.PricingResponse, error) {

	if req.WeightKG <= 0 {
		return nil, errors.New("invalid weight")
	}

	volumetric :=
		(req.LengthCM * req.WidthCM * req.HeightCM) / 6000

	chargeable := math.Max(req.WeightKG, volumetric)

	baseRate := getBaseRate(req.ServiceType)

	baseTariff := chargeable * baseRate

	insurance := 0.0

	if req.UseInsurance {
		insurance = baseTariff * 0.02
	}

	discount := getDiscount(req.PromoCode, baseTariff)

	total := baseTariff + insurance - discount

	return &domain.PricingResponse{
		BaseTariff:   round(baseTariff),
		InsuranceFee: round(insurance),
		Discount:     round(discount),
		TotalPayment: round(total),
		EstimatedSLA: getSLA(req.ServiceType),
	}, nil
}

func getBaseRate(service string) float64 {

	switch strings.ToUpper(service) {

	case "REGULER":
		return 9000

	case "EXPRESS":
		return 18000

	case "SAME_DAY":
		return 30000

	default:
		return 10000
	}
}

func getDiscount(code string, amount float64) float64 {

	switch strings.ToUpper(code) {

	case "HEMAT10":
		return amount * 0.1

	default:
		return 0
	}
}

func getSLA(service string) string {

	switch strings.ToUpper(service) {

	case "REGULER":
		return "3-5 Hari"

	case "EXPRESS":
		return "1-2 Hari"

	case "SAME_DAY":
		return "Hari Yang Sama"

	default:
		return "Unknown"
	}
}

func round(v float64) float64 {
	return math.Round(v*100) / 100
}