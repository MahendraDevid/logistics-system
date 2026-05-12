// tests/functional/warehouse_functional_test.go

//go:build functional
// +build functional

// Tag di atas memastikan test ini HANYA dijalankan kalau eksplisit:
// go test -tags=functional ./tests/functional/...
// Tanpa tag itu, test ini akan diskip otomatis

package functional_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"warehouse-service/internal/domain"
	"warehouse-service/internal/repository"
	"warehouse-service/internal/service"
)

// WarehouseFunctionalSuite adalah test suite untuk functional test
// Suite pattern memudahkan setup dan teardown database
type WarehouseFunctionalSuite struct {
	suite.Suite
	db      *gorm.DB
	service *service.WarehouseService
	ctx     context.Context
}

// SetupSuite dijalankan SEKALI sebelum semua test dalam suite
func (s *WarehouseFunctionalSuite) SetupSuite() {
	// Baca database URL dari environment variable
	// Saat CI/CD, ini akan diisi oleh Docker Compose
	dbURL := os.Getenv("TEST_DATABASE_URL")
	if dbURL == "" {
		dbURL = "host=localhost user=testuser password=testpass dbname=wms_test port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	s.Require().NoError(err, "Gagal koneksi ke database test")

	// PERBAIKAN DI SINI: Auto migrate — Manifest dibuat LEBIH DULU sebelum Package
	err = db.AutoMigrate(&domain.Manifest{}, &domain.Package{})
	s.Require().NoError(err)

	s.db = db
	s.ctx = context.Background()

	// Gunakan repository NYATA (bukan mock)
	repo := repository.NewWarehouseRepository(db)
	// Untuk Kafka, kita masih bisa mock karena Kafka mungkin tidak ada di test environment
	// Atau gunakan embedded Kafka jika diperlukan
	kafkaMock := &NoOpKafkaProducer{} // Simple no-op implementasi
	s.service = service.NewWarehouseService(repo, kafkaMock)
}

// NoOpKafkaProducer adalah implementasi Kafka yang tidak melakukan apa-apa
// Digunakan di functional test agar tidak perlu Kafka nyata
type NoOpKafkaProducer struct{}

func (k *NoOpKafkaProducer) PublishEvent(ctx context.Context, topic, key string, value []byte) error {
	return nil // Diabaikan
}

// SetupTest dijalankan sebelum SETIAP test — bersihkan data
func (s *WarehouseFunctionalSuite) SetupTest() {
	// Hapus semua data test agar setiap test mulai dari kondisi bersih
	s.db.Exec("DELETE FROM packages")
	s.db.Exec("DELETE FROM manifests")
}

// TearDownSuite dijalankan SEKALI setelah semua test selesai
func (s *WarehouseFunctionalSuite) TearDownSuite() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()
}

// TestInboundFlow menguji alur lengkap inbound paket dengan database nyata
func (s *WarehouseFunctionalSuite) TestInboundFlow() {
	awb := "FUNC-AWB-001"
	hubID := "HUB-BANDUNG-01"

	// Jalankan proses inbound
	pkg, err := s.service.ProcessInbound(s.ctx, awb, hubID)

	s.NoError(err)
	s.NotNil(pkg)
	s.Equal(awb, pkg.AWB)
	s.Equal("INBOUND", pkg.Status)

	// Verifikasi data benar-benar tersimpan di database
	var savedPkg domain.Package
	result := s.db.Where("awb = ?", awb).First(&savedPkg)
	s.NoError(result.Error)
	s.Equal(awb, savedPkg.AWB)
	s.Equal("INBOUND", savedPkg.Status)
}

// TestDuplicateInbound memastikan duplikasi AWB ditolak di level DB juga
func (s *WarehouseFunctionalSuite) TestDuplicateInbound() {
	awb := "FUNC-AWB-DUPLIKAT"

	// Inbound pertama — harus sukses
	_, err := s.service.ProcessInbound(s.ctx, awb, "HUB-01")
	s.NoError(err)

	// Inbound kedua — harus gagal
	_, err = s.service.ProcessInbound(s.ctx, awb, "HUB-01")
	s.Error(err)
}

// Entry point untuk menjalankan suite
func TestWarehouseFunctionalSuite(t *testing.T) {
	suite.Run(t, new(WarehouseFunctionalSuite))
}