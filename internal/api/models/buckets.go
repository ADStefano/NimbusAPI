package models

type Buckets struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;unique;not null"`
	BucketName string `gorm:"unique;not null;size:64" json:"bucket_name" binding:"required"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
	UpdatedAt  int64  `gorm:"autoUpdateTime"`
	CreatedBy  string `gorm:"not null;size:32"`
	UpdatedBy  string `gorm:"not null;size:32"`
}
