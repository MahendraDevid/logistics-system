package domain

import (
	"context"
	"time"
)

// CommissionLog mencatat setiap komisi yang didapat kurir
type CommissionLog struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	CourierID   string    `json:"courier_id"`
	AWB         string    `json:"awb"`
	Amount      float64   `json:"amount"`
	Status      string    `json:"status"` // PENDING, PAID
	DeliveredAt time.Time `json:"delivered_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// CourierSummary adalah ringkasan saldo komisi kurir
type CourierSummary struct {
	CourierID       string  `json:"courier_id"`
	TotalDeliveries int     `json:"total_deliveries"`
	TotalAmount     float64 `json:"total_amount"`
	PendingAmount   float64 `json:"pending_amount"`
	PaidAmount      float64 `json:"paid_amount"`
}

// SettlementRepository adalah kontrak akses data
type SettlementRepository interface {
	CreateCommissionLog(ctx context.Context, log *CommissionLog) error
	GetCommissionsByCourier(ctx context.Context, courierID string) ([]CommissionLog, error)
	GetCourierSummary(ctx context.Context, courierID string) (*CourierSummary, error)
	MarkAsPaid(ctx context.Context, courierID string) error
}

// PricingServiceClient adalah kontrak untuk memanggil Pricing Service
// Kita mock ini agar Settlement tidak perlu HTTP call nyata saat test
type PricingServiceClient interface {
	GetCommissionRate(ctx context.Context, serviceType string) (float64, error)
}
