package models
import (
	"time"

)

type Kategoris struct {
	ID            uint `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	User_id string `json:"user_id"`
	Nama_kategori string `json:"nm_kategori"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
