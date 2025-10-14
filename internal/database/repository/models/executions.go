package models

import "database/sql/driver"

type status string

const (
	PENDING     status = "PENDING"
	IN_PROGRESS status = "IN_PROGRESS"
	COMPLETED   status = "COMPLETED"
	FAILED      status = "FAILED"
)

func (s *status) Scan(value interface{}) error {
	*s = status(value.([]byte))
	return nil
}

func (s status) Value() (driver.Value, error) {
	return string(s), nil
}

// Necesário criar o status antes no banco de dados
type Executions struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;unique;not null"`
	BucketID   uint   `gorm:"not null;index;"`
	ObjectID   uint   `gorm:"not null;index;foreignKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status     status `gorm:"type:status;not null;default:'PENDING'"`
	StartedAt  int64  `gorm:"autoCreateTime"`
	EndedAt    int64  `gorm:""`
	CreatedAt  int64  `gorm:"autoCreateTime"`
	UpdatedAt  int64  `gorm:"autoUpdateTime"`
	ExecutedBy string `gorm:"not null;size:32"`
	Success    bool   `gorm:"not null;default:false"`
}

func (Executions) TableName() string {
    return "executions"
}
