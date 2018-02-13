package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"hello/models"
	_ "github.com/go-sql-driver/mysql"
)



func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.200.111:3899)/alw_order?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai", 30)//设置conn中的数据库为默认使用数据库
	// orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.200.111:3899)/alw_order?charset=utf8")
	orm.RegisterDataBase("main", "mysql", "root:123456@/fly?charset=utf8")

	//注册表studentinfo 如果没有会自动创建
	orm.RegisterModel(new(models.User),new(models.Orders))
	//orm.RegisterModel(new(User))//注册表studentinfo 如果没有会自动创建
	//orm.RegisterDriver("mysql", orm.DR_MySQL) //注册mysql驱动
	// orm.RunSyncdb("default",false,true)//后一个使用true会带上很多打印信息，数据库操作和建表操作的；第二个为true代表强制创建表
}


func main() {
	orm.Debug = true //true 打印数据库操作日志信息
	beego.Run()
}

