package service_test

import (
	"context"
	"testing"
	"epod-service/internal/domain"
	"epod-service/internal/service"
	"epod-service/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEPODService_UploadProof_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. Inisialisasi Mocks
	mockRepo := mocks.NewMockEPODRepository(ctrl)
	mockStorage := mocks.NewMockStorageClient(ctrl)
	mockKafka := mocks.NewMockKafkaPublisher(ctrl)

	// 2. Setup ekspektasi
	// Pura-puranya upload S3 sukses dan mengembalikan URL
	mockStorage.EXPECT().
		UploadFile(gomock.Any(), gomock.Any(), "AWB123.jpg").
		Return("https://s3.dummy.com/AWB123.jpg", nil).Times(1)

	// Pura-puranya simpan metadata ke DB sukses
	mockRepo.EXPECT().
		SaveMetadata(gomock.Any(), gomock.Any()).
		Return(nil).Times(1)

	// Pura-puranya publish event Kafka sukses
	mockKafka.EXPECT().
		PublishEvent("PackageDelivered", gomock.Any()).
		Return(nil).Times(1)

	// 3. Inisialisasi Service
	epodSvc := service.NewEPODService(mockRepo, mockStorage, mockKafka)

	// 4. Eksekusi
	req := domain.UploadRequest{
		AWB:         "AWB123",
		CourierID:   "C-999",
		FileBytes:   []byte("dummy-image-data"),
		ContentType: "image/jpeg",
	}
	
	resp, err := epodSvc.ProcessUpload(context.Background(), req)

	// 5. Verifikasi
	assert.NoError(t, err)
	assert.Equal(t, "SUCCESS", resp.Status)
	assert.Equal(t, "https://s3.dummy.com/AWB123.jpg", resp.ImageURL)
}