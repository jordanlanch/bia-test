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
	if startStr == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing end_date parameter"})
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
	consumptions, err := cc.ConsumptionUsecase.GetConsumptionsByPeriod(kind, startStr, endStr, meterIDs, pagination)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, consumptions)
}
