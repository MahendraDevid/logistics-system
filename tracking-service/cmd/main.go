// cmd/main.go
// Entry point aplikasi Tracking & Status Service.
// Menginisialisasi semua dependency dan memulai HTTP server.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tracking-service/internal/cache"
	"tracking-service/internal/handler"
	"tracking-service/internal/kafka"
	"tracking-service/internal/repository"
	"tracking-service/internal/service"
)

func main() {
	// =========================================================
	// Konfigurasi dari Environment Variable
	// =========================================================
	mongoURI  := getEnv("MONGO_URI", "mongodb://localhost:27017")
	mongoDB   := getEnv("MONGO_DB", "tracking_db")
	redisAddr := getEnv("REDIS_ADDR", "localhost:6379")
	redisPass := getEnv("REDIS_PASSWORD", "")
	kafkaBroker := getEnv("KAFKA_BROKER", "localhost:9092")
	appPort  := getEnv("APP_PORT", "8080")

	// =========================================================
	// Inisialisasi MongoDB
	// =========================================================
	db, err := repository.ConnectMongoDB(mongoURI, mongoDB)
	if err != nil {
		log.Fatalf("Gagal koneksi ke MongoDB: %v", err)
	}
	log.Printf("✅ Terhubung ke MongoDB: %s/%s", mongoURI, mongoDB)

	// =========================================================
	// Inisialisasi Redis
	// =========================================================
	redisClient, err := cache.ConnectRedis(redisAddr, redisPass, 0)
	if err != nil {
		log.Fatalf("Gagal koneksi ke Redis: %v", err)
	}
	log.Printf("✅ Terhubung ke Redis: %s", redisAddr)

	// =========================================================
	// Inisialisasi semua layer (Dependency Injection)
	// =========================================================
	repo := repository.NewMongoTrackingRepository(db)

	// Buat index MongoDB saat startup
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := repo.EnsureIndexes(ctx); err != nil {
		log.Printf("WARNING: Gagal buat MongoDB index: %v", err)
	}

	redisCache := cache.NewRedisTrackingCache(redisClient)
	kafkaProducer := kafka.NewTrackingKafkaProducer(kafkaBroker)

	svc := service.NewTrackingService(repo, redisCache, kafkaProducer)
	h := handler.NewTrackingHandler(svc)

	// =========================================================
	// Setup HTTP Routes
	// =========================================================
	mux := http.NewServeMux()

	// Health & Readiness Probes (untuk Kubernetes)
	mux.HandleFunc("/health", h.HealthCheck)
	mux.HandleFunc("/ready", h.HealthCheck)

	// API Endpoints
	// POST /api/v1/tracking/events — Catat event tracking baru
	mux.HandleFunc("/api/v1/tracking/events", h.RecordEvent)

	// GET /api/v1/tracking/{awb}/history — Riwayat perjalanan paket
	// GET /api/v1/tracking/{awb}/status  — Status terakhir paket
	mux.HandleFunc("/api/v1/tracking/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path[len(r.URL.Path)-7:] == "/history" {
			h.GetTrackingHistory(w, r)
		} else if r.URL.Path[len(r.URL.Path)-7:] == "/status" {
			h.GetCurrentStatus(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// =========================================================
	// Start HTTP Server dengan Graceful Shutdown
	// =========================================================
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", appPort),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Goroutine untuk listen shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("🚀 Tracking Service berjalan di port %s", appPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Tunggu signal shutdown
	<-quit
	log.Println("⏳ Menerima signal shutdown, menutup server...")

	// Graceful shutdown: tunggu request yang sedang berjalan (max 30 detik)
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Gagal graceful shutdown: %v", err)
	}

	log.Println("✅ Server berhasil dimatikan")
}

// getEnv mengambil nilai environment variable, atau default jika tidak ada
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
