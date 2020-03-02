package v1

import (
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 获得文章列表
func GetArticleList(c *gin.Context) {
	var article model.Article
	articleList,err:=article.GetList()
	if err!=nil{
		fmt.Printf("错误：%v\n",err.Error())
	}
	c.JSON(200,gin.H{
		"data":articleList,
	})
}

// 获得单个文章
func GetArticle(c *gin.Context) {
	id:=c.Param("id")
	a:=model.Article{}
	i,_:=strconv.Atoi(id)
	a.ID=uint(i)
	article,_:=a.GetOne()
	c.JSON(200,gin.H{
		"data":article,
	})
}

// 新建单个文章
func AddArticle(c *gin.Context) {
	title:=c.PostForm("title")
	author:= c.PostForm("author")
	desc:=c.PostForm("desc")
	content:= c.PostForm("content")
	a:=model.Article{
		Title:       title,
		Author:      author,
		Dsec:        desc,
		Content:     content,
	}
	err:=a.Add()
	if err!=nil{
		c.JSON(404,gin.H{
			"message":"添加失败",
		})
	}

	c.JSON(200,gin.H{
		"msg":"添加成功",
		"data":a,
	})
}

// 修改单个文章
func UpdateArticle(c *gin.Context) {
	
}

// 删除单个文章
func DeleteArticle(c *gin.Context)  {
	
}