package models
import(
	"time"
	"gorm.io/gorm"
)

type Todo struct {
   ID        int64          `json:"id" gorm:"column:id;primarykey"`
   UserId    int64          `json:"user_id" gorm:"column:user_id"`
   Content   string         `json:"content" gorm:"column:content"`
   Completed bool           `json:"completed" gorm:"column:completed;default:false"`
   CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
   UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
   DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index"` // 不需要序列化出去

}