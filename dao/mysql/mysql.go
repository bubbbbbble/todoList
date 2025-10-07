package mysql
import (
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
  	"github.com/spf13/viper"
  	"fmt"

)

var DB *gorm.DB
func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	viper.GetString("mysql.user"),
	viper.GetString("mysql.password"),
	viper.GetString("mysql.addr"),
	viper.GetInt("mysql.port"),
	viper.GetString("mysql.dbname"))
	
	var err error
	DB,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect failed")

	}

	//db.AutoMigrate(&model.Todo{})

}	