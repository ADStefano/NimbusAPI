package models

import "time"

type Buckets struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;unique;not null"`
	BucketName string `gorm:"unique;not null;size:64" json:"bucket_name" binding:"required"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime"`
	CreatedBy  string `gorm:"not null;size:32"`
	UpdatedBy  string `gorm:"not null;size:32"`
}
