package v1

import (
	"blog/model"
	"blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 浏览量+1
func Praise(c *gin.Context) {
	what:=c.Param("what")
	id:=c.Param("id")

	g:=utils.Gin{C:c}

	if what == "article"{
		a:=&model.Article{}
		i,_:=strconv.Atoi(id)
		a.ID=uint(i)

		praisecount,err:=a.Praise()
		if err!=nil{
			g.Response(400,4001,"fail",nil)
		}

		m:=map[string]uint{
			"id":a.ID,
			"praiseCount":praisecount,
		}
		g.Response(200,20001,"success",m)

	}
}
