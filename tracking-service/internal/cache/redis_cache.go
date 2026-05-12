// internal/cache/redis_cache.go
// Implementasi konkret TrackingCache menggunakan Redis.
// Menyimpan status terakhir paket untuk query cepat (O(1)).

package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"tracking-service/internal/domain"

	"github.com/redis/go-redis/v9"
)

const (
	// TTL status di Redis: 7 hari
	// Setelah paket delivered, status akan otomatis expired dari Redis
	statusTTL = 7 * 24 * time.Hour

	// Prefix key di Redis untuk namespace isolation
	keyPrefix = "tracking:status:"
)

// RedisTrackingCache adalah implementasi konkret dari domain.TrackingCache
type RedisTrackingCache struct {
	client *redis.Client
}

// NewRedisTrackingCache membuat instance cache baru
func NewRedisTrackingCache(client *redis.Client) *RedisTrackingCache {
	return &RedisTrackingCache{client: client}
}

// SetStatus menyimpan status terakhir paket ke Redis.
// Key format: "tracking:status:{awb}"
func (c *RedisTrackingCache) SetStatus(ctx context.Context, status *domain.TrackingStatus) error {
	// TODO: implementasi Set ke Redis
	//
	// key := keyPrefix + status.AWB
	// data, err := json.Marshal(status)
	// if err != nil {
	//     return fmt.Errorf("gagal marshal status: %w", err)
	// }
	// if err := c.client.Set(ctx, key, data, statusTTL).Err(); err != nil {
	//     return fmt.Errorf("redis Set gagal: %w", err)
	// }
	// return nil

	_ = json.Marshal // suppress unused import
	_ = keyPrefix
	return fmt.Errorf("not implemented") // placeholder
}

// GetStatus mengambil status terakhir paket dari Redis.
// Mengembalikan nil, nil jika key tidak ditemukan (cache miss).
func (c *RedisTrackingCache) GetStatus(ctx context.Context, awb string) (*domain.TrackingStatus, error) {
	// TODO: implementasi Get dari Redis
	//
	// key := keyPrefix + awb
	// data, err := c.client.Get(ctx, key).Bytes()
	// if err == redis.Nil {
	//     return nil, nil // cache miss — bukan error
	// }
	// if err != nil {
	//     return nil, fmt.Errorf("redis Get gagal: %w", err)
	// }
	// var status domain.TrackingStatus
	// if err := json.Unmarshal(data, &status); err != nil {
	//     return nil, fmt.Errorf("gagal unmarshal status: %w", err)
	// }
	// return &status, nil

	return nil, fmt.Errorf("not implemented") // placeholder
}

// DeleteStatus menghapus status dari Redis (saat paket delivered/returned).
func (c *RedisTrackingCache) DeleteStatus(ctx context.Context, awb string) error {
	// TODO: implementasi Del dari Redis
	//
	// key := keyPrefix + awb
	// if err := c.client.Del(ctx, key).Err(); err != nil {
	//     return fmt.Errorf("redis Del gagal: %w", err)
	// }
	// return nil

	return fmt.Errorf("not implemented") // placeholder
}

// ============================================================
// Helper untuk koneksi Redis — digunakan di main.go
// ============================================================

// ConnectRedis membuat koneksi ke Redis
func ConnectRedis(addr string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,

		// Connection pool settings untuk high load (5K-10K RPS)
		PoolSize:     50,              // Jumlah koneksi di pool
		MinIdleConns: 10,              // Koneksi idle minimum
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("gagal ping Redis: %w", err)
	}

	return client, nil
}
