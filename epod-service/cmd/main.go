package main

import (
	"log"
	"net/http"
	"epod-service/internal/handler"
	"epod-service/internal/kafka"
	"epod-service/internal/repository"
	"epod-service/internal/service"
)

func main() {
	// 1. Load Konfigurasi & Koneksi Eksternal
	db := ConnectMySQL()
	s3Client := ConnectCloudStorage() // Setup koneksi AWS S3 / GCS
	kafkaProducer := kafka.NewKafkaProducer("localhost:9092") // Setup Kafka

	// 2. Dependency Injection (Perakitan)
	epodRepo := repository.NewEPODRepository(db)
	
	// Service e-POD butuh 3 komponen pembantu
	epodService := service.NewEPODService(epodRepo, s3Client, kafkaProducer)
	
	epodHandler := handler.NewEPODHandler(epodService)

	// 3. Setup Router 
	router := SetupRouter()
	// Endpoint menerima file multipart/form-data
	router.Post("/api/v1/epod/upload", epodHandler.UploadProof)

	// 4. Jalankan Server
	log.Println("e-POD Service berjalan di port 8081...")
	http.ListenAndServe(":8081", router)
}