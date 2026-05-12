package main

import (
	"log"
	"net/http"
	// Import folder internal milikmu
	"pricing-service/internal/handler"
	"pricing-service/internal/repository"
	"pricing-service/internal/service"
)

func main() {
	// 1. Load Konfigurasi & Koneksi Database
	db := ConnectPostgreSQL() // Fungsi buatanmu untuk konek DB
	redisClient := ConnectRedis() // Fungsi buatanmu untuk konek Redis

	// 2. Dependency Injection (Perakitan)
	// Masukkan koneksi DB ke dalam Repository
	pricingRepo := repository.NewPricingRepository(db, redisClient)
	
	// Masukkan Repository ke dalam Service
	pricingService := service.NewPricingService(pricingRepo)
	
	// Masukkan Service ke dalam Handler (Controller)
	pricingHandler := handler.NewPricingHandler(pricingService)

	// 3. Setup Router (Misal pakai Chi, Mux, atau Gin)
	router := SetupRouter() // Inisialisasi router
	router.Post("/api/v1/calculate-price", pricingHandler.CalculatePrice)

	// 4. Jalankan Server
	log.Println("Pricing Service berjalan di port 8080...")
	http.ListenAndServe(":8080", router)
}