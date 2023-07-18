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
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10")) // Default limit is 10
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid limit"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0")) // Default offset is 0
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid offset"})
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
	id, err := uuid.Parse(c.Param("id"))
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
