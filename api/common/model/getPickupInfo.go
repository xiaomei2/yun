package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PickupInfo struct {
	ID         int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserID     int    `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	Name       string `gorm:"column:name;NOT NULL" json:"name"`
	Phone      string `gorm:"column:phone;NOT NULL" json:"phone"`
	Text       string `gorm:"column:text" json:"text"`
	NewText    string `gorm:"column:new_text" json:"new_text"`
	PickupName string `gorm:"column:pickup_name;NOT NULL" json:"pickup_name"`
	PickupCode string `gorm:"column:pickup_code;NOT NULL" json:"pickup_code"`
	Status     int64  `gorm:"column:status;default:0;NOT NULL" json:"status"`
	CreateTime int64  `gorm:"column:create_time;default:0;NOT NULL" json:"create_time"`
	UpdateTime int64  `gorm:"column:update_time;default:0;NOT NULL" json:"update_time"`
}

func (m *PickupInfo) TableName() string {
	return "pickup_info"
}
func (m *PickupInfo) Create(db *gorm.DB, ctx context.Context) (int, error) {
	m.CreateTime = time.Now().Unix()
	m.Status = 0
	err := db.WithContext(ctx).Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.ID, nil
}
