package main

import (
	"github.com/gin-gonic/gin"
	"goCommunity/dao"
	"goCommunity/service"
	"log"
)

func main() {
	err := dao.StartDataBase()
	if err != nil {
		log.Fatal(err.Error())
	}

	c := gin.Default()
	c.POST("/send/code", service.SendCode)
	c.POST("/register", service.UserRegister)
	c.POST("/login", service.UserLogin)
	c.Run(":8080")
}
