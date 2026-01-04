// cmd/handlers/handleMeasurements.go
package handlers

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MeasurementHandler struct {
	repo *repositories.MeasurementRepository
}

func NewMeasurementHandler(repo *repositories.MeasurementRepository) *MeasurementHandler {
	return &MeasurementHandler{repo: repo}
}

func (h *MeasurementHandler) HandleCreateMeasurement(c echo.Context) error {
	measurement := models.Measurements{}
	if err := c.Bind(&measurement); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	newMeasurement, err := h.repo.CreateMeasurement(measurement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newMeasurement)
}

func (h *MeasurementHandler) HandleUpdateMeasurement(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	measurement := models.Measurements{}
	c.Bind(&measurement)
	updatedMeasurement, err := h.repo.UpdateMeasurement(measurement, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedMeasurement)
}
