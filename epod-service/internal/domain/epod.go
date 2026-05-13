package domain

type UploadResponse struct {
	Status    string  `json:"status"`
	AWB       string  `json:"awb"`
	CourierID string  `json:"courier_id"`
	ImageURL  string  `json:"image_url"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}