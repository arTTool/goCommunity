package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"goCommunity/dao"
	"goCommunity/util"
	"log"
	"net/http"
	"strconv"
	"time"
)

func UserRegister(c *gin.Context) {
	email := c.PostForm("Email")
	password := c.PostForm("Password")
	code := c.PostForm("Code")
	if email == "" || password == "" || code == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "请将邮箱和密码,验证码填写完整",
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
	r, err := dao.RDB.Get(context.Background(), "CODE_"+email).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "验证码获取存储失败(redis)",
		})
	}
	if code != r {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "验证码错误，请重新输入",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "true",
	})
}

func SendCode(c *gin.Context) {
	email := c.PostForm("Email")
	code := util.GetCode()
	if err := dao.RDB.Set(context.Background(), "CODE_"+email, code, time.Second*60).Err(); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "验证码储存失败(redis)",
		})
		return
	}
	err := util.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "验证码发送失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  true,
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
	token, err := util.GetToken(strconv.Itoa(user.ID), email)
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
