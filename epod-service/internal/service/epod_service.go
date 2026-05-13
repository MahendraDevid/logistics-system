package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/google/uuid"

	"epod-service/internal/domain"
	"epod-service/internal/kafka"
	"epod-service/internal/storage"
)

type EPODService struct {
	storage *storage.LocalStorage
	kafka   *kafka.Producer
}

func NewEPODService(
	storage *storage.LocalStorage,
	kafka *kafka.Producer,
) *EPODService {

	return &EPODService{
		storage: storage,
		kafka:   kafka,
	}
}

func (s *EPODService) Upload(
	file multipart.File,
	header *multipart.FileHeader,
	awb string,
	courierID string,
	lat float64,
	lon float64,
) (*domain.UploadResponse, error) {

	if awb == "" {
		return nil, errors.New("awb required")
	}

	if courierID == "" {
		return nil, errors.New("courier id required")
	}

	err := os.MkdirAll("uploads", os.ModePerm)

	if err != nil {
		return nil, err
	}

	filename := fmt.Sprintf(
		"%s-%s",
		uuid.New().String(),
		header.Filename,
	)

	path := "uploads/" + filename

	err = s.storage.SaveFile(file, path)

	if err != nil {
		return nil, err
	}

	imageURL := "http://localhost:8080/uploads/" + filename

	s.kafka.PublishDeliveredEvent(awb)

	return &domain.UploadResponse{
		Status:    "SUCCESS",
		AWB:       awb,
		CourierID: courierID,
		ImageURL:  imageURL,
		Latitude:  lat,
		Longitude: lon,
	}, nil
}