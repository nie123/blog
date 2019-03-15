package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/models"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbprot := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbprot + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(models.User))

}

func TableName(str string) {
	//return beego.AppConfig.String("dbprefix") + str

}
