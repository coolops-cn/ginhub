package models

import (
	"time"

	"github.com/spf13/cast"
)

type BaseModel struct {
	ID uint32 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimeStampFiled struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
