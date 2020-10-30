package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	dirveName := beego.AppConfig.String("driverName")

	orm.RegisterDriver(dirveName, orm.DRMySQL)
	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//
	dbconn := user+":"+pwd+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8"
	err := orm.RegisterDataBase("default", dirveName, dbconn)
	if err != nil {
		fmt.Print("error")
		return
	}
	fmt.Print("success")
}