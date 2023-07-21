package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jordanlanch/bia-test/domain"
)

type MeterController struct {
	MeterUsecase domain.MeterUsecase
}

func (mc *MeterController) Fetch(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid limit parameter"})
		return
	}

	offsetStr := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid offset parameter"})
		return
	}

	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

	meters, err := mc.MeterUsecase.GetManyMeters(pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meters)
}

func (mc *MeterController) Get(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing id parameter"})
		return
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	meter, err := mc.MeterUsecase.GetMeterById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, meter)
}
