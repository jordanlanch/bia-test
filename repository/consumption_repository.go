package repository

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/bia-test/domain"
	"gorm.io/gorm"
)

// ConsumptionRepository es una implementación de la interfaz domain.ConsumptionRepository.
type consumptionRepository struct {
	db    *gorm.DB
	table string
}

// NewConsumptionRepository crea una nueva instancia de ConsumptionRepository.
func NewConsumptionRepository(db *gorm.DB, table string) domain.ConsumptionRepository {
	return &consumptionRepository{db, table}
}

// Save guarda el consumo en la base de datos.
func (r *consumptionRepository) Save(ctx context.Context, consumption *domain.Consumption) error {
	return r.db.WithContext(ctx).Save(consumption).Error
}

// FindByID busca un consumo por su ID.
func (r *consumptionRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Consumption, error) {
	var consumption domain.Consumption
	err := r.db.WithContext(ctx).First(&consumption, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &consumption, nil
}

// FindByPeriod busca consumos por tipo de período, fechas, IDs de medidores y paginación.
func (r *consumptionRepository) FindByPeriod(ctx context.Context, periodType, start, end string, meterIDs []int, pagination *domain.Pagination) (*domain.Response, error) {
	if !isValidPeriodType(periodType) {
		return nil, fmt.Errorf("Invalid period type: %s", periodType)
	}
	return r.findWithCondition(ctx, "meter_id IN (?) AND timestamp BETWEEN ? AND ?", []interface{}{meterIDs, start, end}, pagination, periodType)
}

// findWithCondition es una función auxiliar para buscar consumos con una condición específica.
func (r *consumptionRepository) findWithCondition(ctx context.Context, condition string, args []interface{}, pagination *domain.Pagination, periodType string) (*domain.Response, error) {
	var consumptions []*domain.Consumption

	datePeriodFormat, datePeriodSQL := getDatePeriodFormat(periodType)

	db := applyPagination(r.db, pagination)
	err := db.WithContext(ctx).
		Select(fmt.Sprintf("%s, meter_id, SUM(active) as active, SUM(reactive_inductive) as reactive_inductive, SUM(reactive_capacitive) as reactive_capacitive, SUM(exported) as exported", datePeriodFormat)).
		Where(condition, args...).
		Group(datePeriodSQL + ", meter_id").
		Order("MIN(timestamp)").
		Find(&consumptions).
		Error

	if err != nil {
		return nil, err
	}

	return formatConsumptionData(consumptions), nil
}

// applyPagination aplica la paginación a la consulta de base de datos.
func applyPagination(db *gorm.DB, pagination *domain.Pagination) *gorm.DB {
	if pagination != nil {
		if pagination.Limit != nil {
			db = db.Limit(*pagination.Limit)
		}
		if pagination.Offset != nil {
			db = db.Offset(*pagination.Offset)
		}
	}
	return db
}

// MeterData es una estructura personalizada para almacenar datos del medidor.
type MeterData struct {
	Period string
	domain.Consumption
}

// formatConsumptionData formats the consumption data into the response structure.
func formatConsumptionData(consumptions []*domain.Consumption) *domain.Response {
	response := &domain.Response{
		Period:    make([]string, 0),
		DataGraph: make([]map[string]interface{}, 0),
	}

	meterDataMap := make(map[int][]MeterData)

	// Group the consumption data by meter ID and collect required data
	for _, consumption := range consumptions {
		meterDataMap[consumption.MeterID] = append(meterDataMap[consumption.MeterID], MeterData{
			Period:      consumption.Period,
			Consumption: *consumption,
		})
	}

	// Process meter data for each meter ID and create data graph items
	for meterID, meterDataSlice := range meterDataMap {
		active := make([]float64, len(meterDataSlice))
		reactiveInductive := make([]float64, len(meterDataSlice))
		reactiveCapacitive := make([]float64, len(meterDataSlice))
		exported := make([]float64, len(meterDataSlice))

		for i, meterData := range meterDataSlice {
			active[i] = meterData.Active
			reactiveInductive[i] = meterData.ReactiveInductive
			reactiveCapacitive[i] = meterData.ReactiveCapacitive
			exported[i] = meterData.Exported
		}

		response.DataGraph = append(response.DataGraph, map[string]interface{}{
			"meter_id":            meterID,
			"address":             "Mock address",
			"active":              active,
			"reactive_inductive":  reactiveInductive,
			"reactive_capacitive": reactiveCapacitive,
			"exported":            exported,
		})
	}

	// Get unique and sorted periods
	response.Period = getUniqueSortedPeriods(meterDataMap)

	return response
}

// getUniqueSortedPeriods returns the unique and sorted periods in date order.
func getUniqueSortedPeriods(meterDataMap map[int][]MeterData) []string {
	periodsMap := make(map[string]bool)

	// Collect unique periods from meter data
	for _, meterDataSlice := range meterDataMap {
		for _, meterData := range meterDataSlice {
			periodsMap[meterData.Period] = true
		}
	}

	// Create a slice to store the unique periods
	sortedPeriods := make([]string, 0, len(periodsMap))

	// Add unique periods to the slice
	for period := range periodsMap {
		sortedPeriods = append(sortedPeriods, period)
	}

	// Sort the periods based on the custom comparison function
	sort.Slice(sortedPeriods, func(i, j int) bool {
		dateI := parsePeriodDate(sortedPeriods[i])
		dateJ := parsePeriodDate(sortedPeriods[j])
		return dateI.Before(dateJ)
	})

	return sortedPeriods
}

// parsePeriodDate parses the period string and returns the corresponding date.
func parsePeriodDate(period string) time.Time {
	// Remove leading/trailing spaces and hyphens
	period = strings.Trim(period, " -")

	// Try to parse the period as "Mon DD"
	date, err := time.Parse("Jan 02", period)
	if err == nil {
		return date
	}

	// Try to parse the period as "Mon YYYY"
	date, err = time.Parse("Jan 2006", period)
	if err == nil {
		return date
	}

	// Try to parse the period as "Mon DD - Mon DD"
	date, err = time.Parse("Jan 02", strings.Split(period, " - ")[0])
	if err == nil {
		return date
	}

	// If none of the formats match, return the zero value
	return time.Time{}
}

// getDatePeriodFormat returns the period format for the given period type.
func getDatePeriodFormat(periodType string) (string, string) {
	switch periodType {
	case "daily":
		return "TO_CHAR(timestamp, 'Mon DD') AS period", "TO_CHAR(timestamp, 'Mon DD')"
	case "weekly":
		// PostgreSQL doesn't have a built-in function to get the week of the year, so we create a date range for 7 days instead.
		return "TO_CHAR(DATE_TRUNC('week', timestamp), 'Mon DD') || ' - ' || TO_CHAR(DATE_TRUNC('week', timestamp) + INTERVAL '6 days', 'Mon DD') AS period", "TO_CHAR(DATE_TRUNC('week', timestamp), 'Mon DD') || ' - ' || TO_CHAR(DATE_TRUNC('week', timestamp) + INTERVAL '6 days', 'Mon DD')"
	case "monthly":
		return "TO_CHAR(timestamp, 'Mon YYYY') AS period", "TO_CHAR(timestamp, 'Mon YYYY')"
	default:
		return "", ""
	}
}

// isValidPeriodType checks if the given period type is valid.
func isValidPeriodType(periodType string) bool {
	validPeriodTypes := map[string]bool{
		"daily":   true,
		"weekly":  true,
		"monthly": true,
	}
	return validPeriodTypes[periodType]
}
