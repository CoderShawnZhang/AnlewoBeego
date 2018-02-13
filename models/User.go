package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int 
	Name      string 
	Password  string
	Remark    string
}

func FindUserByUserName(id int) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id", id).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByUserId(id int) (bool,User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id",id).One(&user)
	return err != orm.ErrNoRows,user
}

func UpdateUser(user *User)() {
	o := orm.NewOrm()
	o.Update(user)
}

func InsertUser(user *User) int64 {
	o := orm.NewOrm()
	id,_ := o.Insert(user)
	return id
}