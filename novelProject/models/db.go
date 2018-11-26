package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var O orm.Ormer

func init(){
	dsn:=beego.AppConfig.String("dsn")
	fmt.Println("dsn is :",dsn)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",dsn)
	orm.RegisterModel(new(Book), new(Chapter),new(Undownload))
	orm.Debug = true
	O = orm.NewOrm()
}
