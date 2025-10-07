package main

import (
	"todolist/dao/mysql"
	"todolist/dao/redis"
	"todolist/settings"
)
func main(){
	settings.Init()
	mysql.Init()
	redis.Init()

} 