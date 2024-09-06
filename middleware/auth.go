package middleware

import (
	"github.com/gin-gonic/gin"
	"goCommunity/util"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		UserClaims, err := util.AnalyzeToken(token)
		if err != nil {
			c.Abort() //防止验证token时执行其他程序
			c.JSON(http.StatusOK, gin.H{
				"code": "-1",
				"msg":  "用户认证失败",
			})
			return
		}
		c.Set("User_claims", UserClaims)
		c.Next()
	}
}
