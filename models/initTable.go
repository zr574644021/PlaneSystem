package models

import (
	"github.com/astaxie/beego/orm"
	"airTicket/models/class"
)

func init()  {
	orm.RegisterModel(new(class.Users),new(class.AirTicket),new(class.Flights),new(class.OrderItem),new(class.Admin))
}
