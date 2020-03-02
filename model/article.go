package model

import (
	"blog/init/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title string `gorm:"size:128;not null" json:"title"`
	Author string `gorm:"size:64;not null" json:"author"`
	Dsec string `gorm:"size:256;not null" json:"dsec"`
	Content string `gorm:"type:text;not null" json:"content"`
	ViewCount uint `gorm:"AUTO_INCREMENT;default:0;column:viewcount" json:"viewcount"`
	PraiseCount uint `gorm:"AUTO_INCREMENT;default:0;column:praisecount"  json:"praisecount"`
}

func (a Article) GetList() ([]*Article,error) {
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return nil,err
	}
	var articles []*Article
	err=DB.Find(&articles).Error
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
	err=DB.Create(&a).Error
	if err!=nil{
		return err
	}
	return nil
}

func Migrate(){
	DB,err:=db.NewConnect()
	if err!=nil{
		fmt.Printf("数据库连接失败：%v\n",err.Error())
		return
	}
	err=DB.AutoMigrate(&Article{}).Error
	if err!=nil{
		fmt.Printf("数据库迁移失败：%v\n",err.Error())
		return
	}
	fmt.Println("数据库迁移完成...")
	return
}