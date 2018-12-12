package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, gin.H{"Code": 0, "Msg": "success", "Data": resp})
}

func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"Code": -1, "Msg": err.Error(), "Data": nil})
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
