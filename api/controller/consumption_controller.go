package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jordanlanch/bia-test/domain"
)

type ConsumptionController struct {
	ConsumptionUsecase domain.ConsumptionUsecase
}

func (cc *ConsumptionController) GetConsumptionsByPeriod(c *gin.Context) {
	meterIDsParam := c.Query("meters_ids")

	var meterIDs []uuid.UUID
	if strings.HasPrefix(meterIDsParam, "[") && strings.HasSuffix(meterIDsParam, "]") {
		// Lista de meter_ids
		meterIDsStr := strings.TrimPrefix(strings.TrimSuffix(meterIDsParam, "]"), "[")
		meterIDsArr := strings.Split(meterIDsStr, ",")
		for _, idStr := range meterIDsArr {
			id, err := uuid.Parse(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
				return
			}
			meterIDs = append(meterIDs, id)
		}
	} else {
		// Solo un meter_id
		id, err := uuid.Parse(meterIDsParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			return
		}
		meterIDs = append(meterIDs, id)
	}
	
	start, err := time.Parse(time.RFC3339, c.Query("start_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	end, err := time.Parse(time.RFC3339, c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

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

	kind := c.Query("kind_period")
	var consumptions []*domain.Consumption
	switch kind {
	case "monthly":
		consumptions, err = cc.ConsumptionUsecase.GetMonthlyConsumptions(start, end, meterIDs, pagination,)
	case "weekly":
		consumptions, err = cc.ConsumptionUsecase.GetWeeklyConsumptions(start, end, meterIDs, pagination)
	case "daily":
		consumptions, err = cc.ConsumptionUsecase.GetDailyConsumptions(start, end, meterIDs, pagination)
	default:
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid period kind"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, consumptions)
}
