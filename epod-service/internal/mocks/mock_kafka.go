package mocks

type MockKafka struct {
}

func NewMockKafka() *MockKafka {
	return &MockKafka{}
}

func (m *MockKafka) PublishDeliveredEvent(awb string) {
	// mock do nothing
}