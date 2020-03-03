package v1

import (
	"blog/model"
	"blog/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 获得文章列表
func GetArticleList(c *gin.Context) {
	var article model.Article
	g:=utils.Gin{C:c}
	articleList,err:=article.GetList()
	if err!=nil{
		g.Response(400,4001,err.Error(),nil)
	}
	g.Response(200,2001,"success",articleList)
}

// 获得单个文章
func GetArticle(c *gin.Context) {
	id:=c.Param("id")
	a:=&model.Article{}
	i,_:=strconv.Atoi(id)
	a.ID=uint(i)
	g:=utils.Gin{C:c}
	article,err:=a.GetOne()
	if err !=nil{
		g.Response(400,4001,err.Error(),nil)
	}

	//浏览量+1
	if err=a.View();err!=nil{
		fmt.Println(err.Error())
	}

	g.Response(200,2001,"success",article)
}

// 新建单个文章
func AddArticle(c *gin.Context) {
	title:=c.PostForm("title")
	author:= c.PostForm("author")
	desc:=c.PostForm("desc")
	content:= c.PostForm("content")
	a:=&model.Article{
		Title:       title,
		Author:      author,
		Dsec:        desc,
		Content:     content,
	}

	g:=utils.Gin{C:c}

	err:=a.Add()
	if err!=nil{
		g.Response(500,5001,err.Error(),nil)
	}

	g.Response(200,20001,"success",a)
}

// 修改单个文章
func UpdateArticle(c *gin.Context) {
	
}

// 删除单个文章
func DeleteArticle(c *gin.Context)  {
	
}

