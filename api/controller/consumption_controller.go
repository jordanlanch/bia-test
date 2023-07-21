package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jordanlanch/bia-test/domain"
)

type ConsumptionController struct {
	ConsumptionUsecase domain.ConsumptionUsecase
}

func (cc *ConsumptionController) GetConsumptionsByPeriod(c *gin.Context) {
	meterIDsParam := c.Query("meters_ids")
	if meterIDsParam == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing meters_ids parameter"})
		return
	}

	var meterIDs []int
	if strings.HasPrefix(meterIDsParam, "[") && strings.HasSuffix(meterIDsParam, "]") {
		// Lista de meter_ids
		meterIDsStr := strings.TrimPrefix(strings.TrimSuffix(meterIDsParam, "]"), "[")
		meterIDsArr := strings.Split(meterIDsStr, ",")
		for _, idStr := range meterIDsArr {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
				return
			}
			meterIDs = append(meterIDs, id)
		}
	} else {
		// Solo un meter_id
		id, err := strconv.Atoi(meterIDsParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			return
		}
		meterIDs = append(meterIDs, id)
	}

	startStr := c.Query("start_date")
	if startStr == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing start_date parameter"})
		return
	}

	endStr := c.Query("end_date")
	if endStr == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing end_date parameter"})
		return
	}

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

	kind := c.Query("kind_period")
	consumptions, err := cc.ConsumptionUsecase.GetConsumptionsByPeriod(kind, startStr, endStr, meterIDs, pagination)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, consumptions)
}
