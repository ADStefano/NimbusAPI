package models

type Objects struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;unique;not null"`
	BucketID  uint   `gorm:"not null;index;foreignKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Key       string `gorm:"not null;unique;size:64"`
	Prefix    string `gorm:"size:64"`
	Size      int64  `gorm:"not null"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	CreatedBy string `gorm:"not null;size:32"`
	UpdatedBy string `gorm:"not null;size:32"`
}
