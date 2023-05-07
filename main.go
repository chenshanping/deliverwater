package main

import (
	"deliverwater/initialize"
	"deliverwater/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	initialize.InitConfig()
	initialize.InitDB()
	list, total := service.UserService.UserList(-1, -1)
	a := gin.H{
		"total": total,
		"list":  list,
	}
	fmt.Printf("%v", a)

}
