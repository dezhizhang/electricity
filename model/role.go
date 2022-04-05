package model

type Role struct {
	ID          int32     `gorm:"primaryKey"json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int32     `json:"status"`
	AddTime     int64 `json:"add_time"`
}
