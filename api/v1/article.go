package v1

import (
	"blog/pkg/utils"
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获得文章列表
func GetArticleList(c *gin.Context) {
	g:=utils.Gin{C:c}

	var article service.ArticleGetList

	// 数据绑定
	err:=c.ShouldBind(&article)
	if err!=nil{
		g.Response(400,4001,fmt.Sprintf("failed:%v",err.Error()),nil)
		return
	}

	// 数据验证
	_,errList:=article.Check()
	if errList!=nil{
		g.Response(400,4001,fmt.Sprintf("failed:%v",errList),nil)
		return
	}

	// 数据写入
	articleList,err:=article.Execute()
	if err!=nil{
		g.Response(500,5001,fmt.Sprintf("failed:%v",err.Error()),nil)
		return
	}

	g.Response(200,2001,"success",articleList)
}

// 获得单个文章
func GetArticle(c *gin.Context) {
	idStr:=c.Param("id")
	id:=com.StrTo(idStr).MustInt()

	var article service.ArticleGetOne
	article.ID=uint(id)

	g:=utils.Gin{C:c}

	// 数据验证
	_,errList:=article.Check()
	if errList!=nil{
		g.Response(400,4001,fmt.Sprintf("failed:%v",errList),nil)
		return
	}

	// 数据写入
	a,err:=article.Execute()
	if err !=nil{
		g.Response(400,4001,err.Error(),nil)
		return
	}


	//浏览量+1
	if err=a.View();err!=nil{
		fmt.Println(err.Error())
	}

	g.Response(200,2001,"success",a)
}

// 新建单个文章
func AddArticle(c *gin.Context) {
	g:=utils.Gin{C:c}

	var article service.ArticleAdd

	// 数据绑定
	err:=c.ShouldBind(&article)
	if err!=nil{
		g.Response(400,4001,fmt.Sprintf("failed:%v",err.Error()),nil)
		return
	}

	// 数据验证
	_,errList:=article.Check()
	if errList!=nil{
		g.Response(400,4001,fmt.Sprintf("failed:%v",errList),nil)
		return
	}

	// 数据写入
	a,err:=article.Execute()
	if err!=nil{
		g.Response(500,5001,fmt.Sprintf("failed:%v",err.Error()),nil)
		return
	}

	// 数据返回
	g.Response(200,20001,"success",a)
}

// 修改单个文章
func UpdateArticle(c *gin.Context) {
	idStr:=c.Param("id")
	id:=com.StrTo(idStr).MustInt()

	g:=utils.Gin{C:c}


	// 数据绑定
	var article service.ArticleUpdata
	err:=c.ShouldBind(&article)
	if err!=nil{
		g.Response(400,4001,fmt.Sprintf("failed:%v",err.Error()),nil)
		return
	}
	article.ID=uint(id)

	// 数据验证
	_,errList:=article.Check()
	if errList!=nil{
		g.Response(400,4001,fmt.Sprintf("failed:%v",errList),nil)
		return
	}

	// 数据写入
	a,err:=article.Execute()
	if err!=nil{
		g.Response(500,5001,fmt.Sprintf("failed:%v",err.Error()),nil)
		return
	}

	// 数据返回
	g.Response(200,20001,"success",a)
}

// 删除单个文章
func DeleteArticle(c *gin.Context)  {
	idStr:=c.Param("id")
	id:=com.StrTo(idStr).MustInt()

	g:=utils.Gin{C:c}

	// 数据绑定
	var article service.ArticleDelete
	article.ID=uint(id)

	// 数据验证
	_,errList:=article.Check()
	if errList!=nil{
		g.Response(400,4001,fmt.Sprintf("failed:%v",errList),nil)
		return
	}

	// 数据写入
	err:=article.Execute()
	if err!=nil{
		g.Response(500,5001,fmt.Sprintf("failed:%v",err.Error()),nil)
		return
	}

	// 数据返回
	g.Response(200,20001,"success",article)
}

