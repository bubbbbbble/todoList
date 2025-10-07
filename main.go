package main

import (
	"todolist/dao/mysql"
	"todolist/dao/redis"
	"todolist/routes"
	"todolist/settings"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)
func main(){
	settings.Init()
	mysql.Init()
	redis.Init()
	r := gin.Default()
	routes.InitRouter(r)
	r.Run(":"+ viper.GetString("port"))
} 