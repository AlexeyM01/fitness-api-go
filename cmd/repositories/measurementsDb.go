// fitness-api/cmd/repositories/measurementsDb.go
package repositories

import (
	"database/sql"
	"fitness-api/cmd/models"
	"time"
)

type MeasurementRepository struct {
	db *sql.DB
}

func NewMeasurementRepository(db *sql.DB) *MeasurementRepository {
	return &MeasurementRepository{db: db}
}

func (r *MeasurementRepository) CreateMeasurement(measurement models.Measurements) (models.Measurements, error) {
	sqlStatement := "INSERT INTO measurements (user_id, weight, height, body_fat, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.db.QueryRow(sqlStatement, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
	if err != nil {
		return measurement, err
	}

	return measurement, nil
}

func (r *MeasurementRepository) UpdateMeasurement(measurement models.Measurements, id int) (models.Measurements, error) {
	sqlStatement := `
    UPDATE measurements
    SET weight = $2, height = $3, body_fat = $4, created_at = $5
    WHERE id = $1
    RETURNING id`
	err := r.db.QueryRow(sqlStatement, id, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&id)
	if err != nil {
		return models.Measurements{}, err
	}
	measurement.Id = id
	return measurement, nil
}
