package service_test

import (
	"context"
	"testing"
	"pricing-service/internal/domain"
	"pricing-service/internal/service"
	"pricing-service/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPricingService_CalculateTariff_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. Inisialisasi Mock
	mockRepo := mocks.NewMockPricingRepository(ctrl)

	// 2. Setup ekspektasi (apa yang dikembalikan oleh mock database)
	// Misalnya: Asumsi dari DB tarif dasarnya adalah 10.000 per kg
	mockRepo.EXPECT().
		GetBaseTariff(gomock.Any(), "JKT", "BDG").
		Return(10000.0, nil).Times(1)

	// 3. Inisialisasi Service dengan Mock Repo
	pricingSvc := service.NewPricingService(mockRepo)

	// 4. Jalankan fungsi yang di-test
	req := domain.CalculationRequest{
		Origin:      "JKT",
		Destination: "BDG",
		WeightKG:    2.0,
		ServiceType: "REGULAR",
	}
	
	resp, err := pricingSvc.CalculateTariff(context.Background(), req)

	// 5. Verifikasi hasil
	assert.NoError(t, err)
	assert.Equal(t, 20000.0, resp.TotalTariff) // 2 kg * 10.000
	assert.Equal(t, "2-3 Days", resp.SLA)
}