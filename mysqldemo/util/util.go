package util

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)
var db *sql.DB

func InitMysql()  {
	fmt.Println("Init mysql....")

	dirveName := beego.AppConfig.String("driverName")

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//
	dbconn := user+":"+pwd+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8"
	db1,err := sql.Open( dirveName, dbconn)
	if err != nil {
		fmt.Println(err.Error())

	}else{
		db = db1
		CreateTableWithUser()
	}

}

func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDb(sql)
}

//操作数据库
//insert user values(?,?,?) values (1,2,3)
func ModifyDb(sql string,args ...interface{})(int64,error){
	result,err := db.Exec(sql,args...) //1,2,3
	if err != nil{
		log.Println(err)
		return 0,err
	}
	count,err := result.RowsAffected()
	if err != nil{
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row  {
	return db.QueryRow(sql)
}