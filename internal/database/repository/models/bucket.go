package models

type Bucket struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;unique;not null"`
	BucketName string `gorm:"unique;not null;size:64"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
	UpdatedAt  int64  `gorm:"autoUpdateTime"`
	CreatedBy  string `gorm:"not null;size:32"`
	UpdatedBy  string `gorm:"not null;size:32"`
}