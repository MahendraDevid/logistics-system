package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"epod-service/internal/service"
)

type EPODHandler struct {
	service *service.EPODService
}

func NewEPODHandler(
	service *service.EPODService,
) *EPODHandler {

	return &EPODHandler{
		service: service,
	}
}

func (h *EPODHandler) Upload(c *gin.Context) {

	awb := c.PostForm("awb")
	courierID := c.PostForm("courier_id")

	lat, _ := strconv.ParseFloat(
		c.PostForm("latitude"),
		64,
	)

	lon, _ := strconv.ParseFloat(
		c.PostForm("longitude"),
		64,
	)

	fileHeader, err := c.FormFile("image")

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	file, err := fileHeader.Open()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	defer file.Close()

	result, err := h.service.Upload(
		file,
		fileHeader,
		awb,
		courierID,
		lat,
		lon,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, result)
}