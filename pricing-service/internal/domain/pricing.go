package domain

type CalculationRequest struct {
	Origin      string
	Destination string

	WeightKG float64

	Length float64
	Width  float64
	Height float64

	ServiceType string
}

type CalculationResponse struct {
	BaseTariff float64
	Insurance  float64
	Discount   float64
	Total      float64
	Estimated  string
}