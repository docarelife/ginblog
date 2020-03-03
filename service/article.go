package service

import (
	"blog/model"
	"blog/pkg/utils"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
)

// 获取文章列表
type ArticleGetList struct {
	PageNum  uint `json:"pageNum" form:"pageNum" valid:"Required"`
	PageSize uint `json:"pageSize" form:"pageSize" valid:"Required"`
}
func (agl ArticleGetList) Check() (bool, interface{}) {
	v := validation.Validation{}
	b, err := v.Valid(&agl)
	if err != nil {
		return b, err
	}
	if !b {
		return b, v.Errors
	}
	return b, nil
}
func (agl ArticleGetList) Execute() ([]*model.Article,error) {

	offset,limit:=utils.GetOffset(agl.PageNum,agl.PageSize)

	var a model.Article
	article,err:=a.GetList(offset,limit)
	if err!=nil{
		return nil,err
	}
	return article,nil
}



// 获取文章详情
type ArticleGetOne struct {
	ID      uint   `valid:"Required;Min(1)"`
}
func (ago ArticleGetOne) Check() (bool, interface{}) {
	v := validation.Validation{}
	b, err := v.Valid(&ago)
	if err != nil {
		return b, err
	}
	if !b {
		return b, v.Errors
	}
	return b, nil
}

func (ago ArticleGetOne) Execute() (*model.Article,error) {
	a:=model.Article{
		Model:       gorm.Model{ID: ago.ID},
	}
	article,err:=a.GetOne()
	if err!=nil{
		return nil,err
	}
	return article,nil
}


// 添加文章
type ArticleAdd struct {
	Title   string `json:"title" form:"title" valid:"Required;MinSize(2);MaxSize(128)"`
	Author  string `json:"author" form:"author" valid:"Required;MaxSize(64)"`
	Dsec    string `json:"dsec" form:"desc" valid:"Required;MaxSize(256)"`
	Content string `json:"content" form:"content" valid:"Required"`
}

func (aa ArticleAdd) Check() (bool, interface{}) {
	v := validation.Validation{}
	b, err := v.Valid(&aa)
	if err != nil {
		return b, err
	}
	if !b {
		return b, v.Errors
	}
	return b, nil
}
func (aa ArticleAdd) Execute() (*model.Article,error) {
	a:=model.Article{
		Title:       aa.Title,
		Author:      aa.Author,
		Dsec:        aa.Dsec,
		Content:     aa.Content,
	}
	err:=a.Add()
	if err!=nil{
		return nil,err
	}
	return &a,nil
}

// 更新文章
type ArticleUpdata struct {
	ID      uint   `valid:"Required"`
	Title   string `json:"title" form:"title" valid:"Required;MinSize(2);MaxSize(128)"`
	Author  string `json:"author" form:"author" valid:"Required;MaxSize(64)"`
	Dsec    string `json:"dsec" form:"desc" valid:"Required;MaxSize(256)"`
	Content string `json:"content" form:"content" valid:"Required"`
}

func (au ArticleUpdata) Check() (bool, interface{}) {
	v := validation.Validation{}
	b, err := v.Valid(&au)
	if err != nil {
		return b, err
	}
	if !b {
		return b, v.Errors
	}
	return b, nil
}

func (au ArticleUpdata) Execute() (*model.Article,error) {
	a:=model.Article{
		Model:       gorm.Model{ID: au.ID},
		Title:       au.Title,
		Author:      au.Author,
		Dsec:        au.Dsec,
		Content:     au.Content,
	}
	err:=a.Update()
	if err!=nil{
		return nil,err
	}
	return &a,nil
}

// 删除文章
type ArticleDelete struct {
	ID uint `valid:"Required"`
}

func (ad ArticleDelete) Check() (bool, interface{}) {
	v := validation.Validation{}
	b, err := v.Valid(&ad)
	if err != nil {
		return b, err
	}
	if !b {
		return b, v.Errors
	}
	return b, nil
}

func (ad ArticleDelete) Execute() error {
	a:=model.Article{
		Model:       gorm.Model{ID: ad.ID},
	}
	err:=a.Delete()
	if err!=nil{
		return err
	}
	return nil
}