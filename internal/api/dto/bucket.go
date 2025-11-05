package dto

import "time"

type BucketRequest struct {
	Name      string `json:"name" binding:"required"`
	CreatedBy string `json:"created_by"`
}

type BucketResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}
