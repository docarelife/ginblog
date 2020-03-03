package utils

import "github.com/gin-gonic/gin"

func GetIP(c *gin.Context) string {
	ipStr:=c.ClientIP()
	return ipStr
}
