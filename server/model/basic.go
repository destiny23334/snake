package model

import (
	"gorm.io/gorm"
	"time"
)

// Basic 字段是每个数据库表都需要包含的
type Basic struct {
	gorm.Model
	ID        uint       `gorm:"primarykey;autoIncrement" json:"id"`
	CreatedAt time.Time  `gorm:"type:timestamp(0)" json:"created_at"`
	DeletedAt *time.Time `gorm:"type:timestamp(0)" json:"deleted_at"` // 删除时间默认为空，所以设置为了指针
	UpdatedAt time.Time  `gorm:"type:timestamp(0)" json:"updated_at"`
}

// BeforeCreate 自动更新创建时间
func (b *Basic) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreatedAt = time.Now()
	return
}

// BeforeUpdate 自动更新更新时间
func (b *Basic) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now()
	return
}

// BeforeDelete 自动更新删除时间
func (b *Basic) BeforeDelete(tx *gorm.DB) (err error) {
	if b.DeletedAt == nil {
		now := time.Now()
		b.DeletedAt = &now
	}
	return
}
