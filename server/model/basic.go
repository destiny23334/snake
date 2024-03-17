package model

import "time"

// Basic 字段是每个数据库表都需要包含的
type Basic struct {
	ID        uint      `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	DeletedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
}
