package service

import (
	"github.com/gin-gonic/gin"
	"goCommunity/dao"
	"goCommunity/util"
	"net/http"
)

func UserRegister(c *gin.Context) {
	email := c.PostForm("Email")
	password := c.PostForm("Password")
	if email == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "请将邮箱和密码填写完整",
		})
		return
	}
	err := dao.UserRegister(email, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "创建用户失败，请联系管理员",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "true",
	})
}

func UserLogin(c *gin.Context) {
	email := c.PostForm("Email")
	password := c.PostForm("Password")
	if email == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "邮箱或密码不能为空",
		})
		return
	}
	user, err := dao.GetUser(email, util.GetMd5(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  false,
			"data": gin.H{
				"err": "[GET DATABASE ERROR]:" + err.Error(),
			},
		})
		return
	}
	token, err := util.GetToken(user.ID, email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "true",
		"data": gin.H{
			"token": token,
		},
	})
}
