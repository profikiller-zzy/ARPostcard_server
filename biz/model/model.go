package model

type MODEL struct {
	ID        int   `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt int64 `gorm:"column:created_at" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" json:"updated_at"`
}
