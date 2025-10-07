package models
import (
	"time"
	"gorm.io/gorm"
)
type User struct {
   ID        int64          `json:"id" gorm:"column:id;primarykey"`
   Username  string         `json:"username" gorm:"column:username;unique"`
   Password  string         `json:"-" gorm:"column:password"`
   Email     string         `json:"email" gorm:"column:email;unique"`
   Avatar    string         `json:"avatar" gorm:"column:avatar"`
   CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
   UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
   DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"` // 不需要序列化出去
}

