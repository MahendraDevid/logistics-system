package mocks

import "github.com/golang/mock/gomock"

type MockPricingRepository struct {
	rate float64
}

type MockPricingRepositoryRecorder struct {
	mock *MockPricingRepository
}

func NewMockPricingRepository(
	ctrl *gomock.Controller,
) *MockPricingRepository {

	return &MockPricingRepository{
		rate: 10000,
	}
}

func (m *MockPricingRepository) EXPECT() *MockPricingRepositoryRecorder {
	return &MockPricingRepositoryRecorder{
		mock: m,
	}
}

func (r *MockPricingRepositoryRecorder) GetBaseRate(
	serviceType interface{},
) *MockPricingRepositoryRecorder {

	return r
}

func (r *MockPricingRepositoryRecorder) Return(
	rate float64,
) {

	r.mock.rate = rate
}

func (m *MockPricingRepository) GetBaseRate(
	serviceType string,
) float64 {

	return m.rate
}