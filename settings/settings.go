package settings
import (
	"github.com/spf13/viper"
)

//获取配置文件
func Init() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath("./settings")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
