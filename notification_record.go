package root

import (
	"time"
)

type NotificationRecord struct {
	ID           uint64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	GroupID      uint64     `gorm:"column:group_id;not null" json:"group_id"`
	Notification string     `gorm:"column:notification;not null" json:"notification"`
	CreatedAt    time.Time  `gorm:"column:created_at;not null;DEFAULT:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;not null;DEFAULT:CURRENT_TIMESTAMP"json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type NotificationRecordService interface {
	Create(user *GroupUserRelation) error
	Update(user *GroupUserRelation, fields map[string]interface{}) error
	Find(id uint64) (*GroupUserRelation, error)
	FindByName(name string) (*GroupUserRelation, error)
	Delete(user *GroupUserRelation) error
}
