package utils

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Base `json:"baseInfo"`
}

type Base struct {
	YourIP string `json:"yourIP"`
	From string `json:"from"`
}

func (g Gin) Response(httpcode int,code int,msg string,data interface{}) {
	var r Response
	r.YourIP=GetIP(g.C)
	r.From="www.gpdream.com"
	r.StatusCode=code
	r.Message=msg
	r.Data=data

	g.C.JSON(httpcode,r)
}



