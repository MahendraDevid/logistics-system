package mocks

import "mime/multipart"

type MockStorage struct {
}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func (m *MockStorage) SaveFile(
	file multipart.File,
	path string,
) error {

	return nil
}