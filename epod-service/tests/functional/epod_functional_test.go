package functional_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestUploadEPOD_Functional(t *testing.T) {
	// Setup router test
	// router := setupTestRouter()

	// Membuat body multipart/form-data
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	
	// Simulasi field teks
	writer.WriteField("awb", "AWB-TEST-001")
	writer.WriteField("courier_id", "C-001")
	writer.WriteField("gps_lat", "-6.200000")
	writer.WriteField("gps_long", "106.816666")

	// Simulasi file gambar (dummy)
	part, _ := writer.CreateFormFile("image", "test_photo.jpg")
	part.Write([]byte("fake-image-content"))
	writer.Close()

	// Buat HTTP request
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/epod/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	
	// Eksekusi
	// router.ServeHTTP(w, req)

	// Verifikasi status dan respons JSON
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "SUCCESS")
}