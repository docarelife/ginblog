package model

import (
	"blog/init/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title string `gorm:"size:128;not null" json:"title" form:"title" valid:"Required;MinSize(2);MaxSize(128)"`
	Author string `gorm:"size:64;not null" json:"author" form:"author" valid:"Required;MaxSize(64)"`
	Dsec string `gorm:"size:256;not null" json:"dsec" form:"desc" valid:"Required;MaxSize(256)"`
	Content string `gorm:"type:text;not null" json:"content" form:"content" valid:"Required"`
	ViewCount uint `gorm:"AUTO_INCREMENT;default:0;column:viewcount" json:"viewcount"`
	PraiseCount uint `gorm:"AUTO_INCREMENT;default:0;column:praisecount"  json:"praisecount"`
}

func (a Article) GetList(offset uint,size uint) ([]*Article,error) {
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return nil,err
	}
	defer DB.Close()

	var articles []*Article
	err=DB.Offset(offset).Limit(size).Find(&articles).Error
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return nil,err
	}

	return articles,nil
}

func (a Article) GetOne() (*Article,error) {
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return nil,err
	}
	defer DB.Close()

	var article Article

	err=DB.Where("id=?",a.ID).First(&article).Error
	if err!=nil{
		return nil,err
	}
	return &article,nil
}

func (a Article) Add() error {
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return err
	}
	defer DB.Close()

	err=DB.Create(&a).Error
	if err!=nil{
		return err
	}
	return nil
}

func (a Article) Update() error {
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return err
	}
	defer DB.Close()

	err=DB.Save(&a).Error
	if err!=nil{
		return err
	}
	return nil
}

func (a Article) Delete() error {
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return err
	}
	defer DB.Close()

	err=DB.Delete(&a).Error
	if err!=nil{
		return err
	}
	return nil
}

func (a Article) View() error {
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return err
	}
	defer DB.Close()

	err=DB.Model(&a).UpdateColumn("viewcount",gorm.Expr("viewcount + ?",1)).Error
	if err !=nil {
		return err
	}
	return nil
}

func (a Article) Praise() (uint,error) {
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return 0,err
	}
	defer DB.Close()

	err=DB.Model(&a).UpdateColumn("praisecount",gorm.Expr("praisecount + ?",1)).Error
	if err !=nil {
		return 0,err
	}

	var article Article
	DB.Where("id = ?",a.ID).Select("praisecount").First(&article)
	return article.PraiseCount,nil
}


func Migrate(){
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return
	}
	defer DB.Close()

	err=DB.AutoMigrate(&Article{}).Error
	if err!=nil{
		fmt.Printf("数据库迁移失败：%v\n",err.Error())
		return
	}
	fmt.Println("数据库迁移完成...")
	return
}