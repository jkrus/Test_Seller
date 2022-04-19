package models

import (
	"time"
)

type (
	Announcement struct {
		ID          int64     `json:"id,omitempty" gorm:"id, index, primary-key"`
		Name        string    `json:"name" gorm:"name"`
		Description string    `json:"description,omitempty" gorm:"description"`
		Price       float64   `json:"price" gorm:"price"`
		Images      []byte    `json:"images,omitempty" gorm:"images"`
		MainImage   string    `json:"mainImage" gorm:"main_image"`
		CreatedAt   time.Time `gorm:"data, default, now()"`
	}

	AnnouncementDTO struct {
		Name        string   `json:"name"`
		Description string   `json:"description,omitempty"`
		Price       float64  `json:"price"`
		Images      []string `json:"images,omitempty"`
	}

	ImagesDTO struct {
		Images map[string]struct{} `json:"images"`
	}
)
