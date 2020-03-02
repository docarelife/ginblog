package db

import (
	"blog/init/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() {
	_,err:=NewConnect()
	if err!=nil{
		panic(fmt.Sprintf("初始化数据库失败：%v\n",err.Error()))
	}
	fmt.Println("初始化数据库完成...")
}

func NewConnect() (*gorm.DB,error){
	c:=config.GlobalConfig.Mysql
	connectStr:=fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",c.Username,c.Password,c.IP,c.Port,c.DBName,c.Param)
	db,err:=gorm.Open("mysql",connectStr)
	if err!=nil{
		fmt.Printf("连接数据库出错：%v\n",err.Error())
		return nil,err
	}
	return db,nil
}

