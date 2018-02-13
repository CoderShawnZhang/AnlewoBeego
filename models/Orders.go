package models

import (
	"github.com/astaxie/beego/orm"
)

type Orders struct {
	Orderid      int `orm:"pk;auto"`
	Ordersn      string
}

func FindOrderById(orderId int) (bool, Orders) {
	o := orm.NewOrm()
	var orders Orders
	err := o.QueryTable(orders).Filter("orderId",orderId).One(&orders)
	return err != orm.ErrNoRows,orders
}