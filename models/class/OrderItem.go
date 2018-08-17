package class

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type OrderItem struct {
	Id 				int
	Name    		string
	Identify		string
	Address 		string
	PhoneNum		string
	BuyStatus		int			`orm:"default(0)"`
	DyStatus		int			`orm:"default(0)"`
	AirTic			*AirTicket	`orm:"rel(fk)"`
	User 			*Users		`orm:"rel(fk)"`
}

func (c *OrderItem)Add()  bool{
	o := orm.NewOrm()
	if _, err := o.Insert(c); err != nil {
		beego.Error("add buy_user error ",err)
		return false
	}
	return true
}

//添加到购物车
func AddBuyCar(buyUser OrderItem) bool {
	o := orm.NewOrm()
	var ticket AirTicket
	if  err := o.QueryTable("air_ticket").Filter("name",buyUser.AirTic.Name).One(&ticket) ; err != nil || ticket.Seats <=0{
		return false
	}else {
		ticket.Seats--
		if  !ticket.UpdateAirTicket() {
			beego.Error("add buyuser ticket error ",err)
			return false
		}else {
			var buyTic  OrderItem
			var user Users
			if err := o.QueryTable("users").Filter("username",buyUser.User.UserName).One(&user); err != nil {
				return false
			}
			buyTic.Name = buyUser.Name
			buyTic.AirTic = buyUser.AirTic
			buyTic.User = &user
			buyTic.Address = buyUser.Address
			buyTic.Identify = buyUser.Identify
			buyTic.PhoneNum = buyUser.PhoneNum
			fmt.Println(buyTic.User.Id)
			if !buyTic.Add() {
				beego.Error("add buyuser error ",err)
				return false
			}
			return true
		}
	}
}

//从购物车删除
func DelBuyCar(id int, ariName string) bool {
	o := orm.NewOrm()

	if _, err := o.QueryTable("order_item").Filter("id",id).Filter("buy_status",0).Delete(); err != nil {
		beego.Error("delete BuyCar error ",err)
		return false
	}else {
		var ticket AirTicket
		if  err := o.QueryTable("air_ticket").Filter("name",ariName).One(&ticket) ; err != nil {
			return false
		}else {
			ticket.Seats++
			if !ticket.UpdateAirTicket() {
				return false
			}
			return true
		}
	}
}

//付款
func BuyCarResult(allItem []OrderItem) bool {
	o := orm.NewOrm()
	var flag bool
	var item OrderItem

	for _, u := range allItem {
		if err := o.QueryTable("order_item").Filter("id",u.Id).Filter("buy_status",0).One(&item); err == nil {
			if _,err := o.QueryTable("order_item").Filter("id",u.Id).Update(orm.Params{"buy_status":1}); err!= nil {
				flag = false
			}else {
				flag = true
			}
		}else {
			flag = false
		}
	}
	return flag
}

//打印
func DYResult(allItem OrderItem) bool {
	o := orm.NewOrm()
	var flag bool
	var item OrderItem

	if err := o.QueryTable("order_item").Filter("id",allItem.Id).Filter("dy_status",0).One(&item); err == nil {
		if _,err := o.QueryTable("order_item").Filter("id",allItem.Id).Update(orm.Params{"dy_status":1}); err!= nil {
			flag = false
		}else {
			flag = true
		}
	}else {
		flag = false
	}
	return flag
}

//获取购物车信息
func GetCarMessage(name string)  (bool, int64, []OrderItem) {
	o := orm.NewOrm()
	var allItem []OrderItem
	var user Users
	if err := o.QueryTable("users").Filter("user_name",name).One(&user); err != nil {
		return false, 0, nil
	}
	if num,err := o.QueryTable("order_item").Filter("user_id",user.Id).Filter("buy_status",0).RelatedSel().All(&allItem); err != nil {
		beego.Error("get all buyuser error ",err)
		return false, num, nil
	}else {
		return true, num, allItem
	}
}

//获取订票信息
func GetOrder(name string)  (bool, int64, []OrderItem) {
	o := orm.NewOrm()
	var allItem []OrderItem
	var user Users
	if err := o.QueryTable("users").Filter("user_name",name).One(&user); err != nil {
		return false, 0, nil
	}
	if num,err := o.QueryTable("order_item").Filter("user_id",user.Id).RelatedSel().All(&allItem); err != nil {
		beego.Error("get all buyuser error ",err)
		return false, num, nil
	}else {
		return true, num, allItem
	}
}

//管理员获取所有购票信息
func AdminGetAllBuyUser() (bool, int64, []OrderItem) {
	o := orm.NewOrm()
	var allItem []OrderItem
	if num,err := o.QueryTable("order_item").RelatedSel().OrderBy("name").All(&allItem); err != nil {
		beego.Error("get all buyuser error ",err)
		return false, num, nil
	}else {
		return true, num, allItem
	}
}

//用户获取所有购票信息
func UserGetAllBuy(name string) (bool, int64, []OrderItem) {
	o := orm.NewOrm()
	var allItem []OrderItem
	var user Users
	if err := o.QueryTable("users").Filter("user_name",name).One(&user); err != nil {
		return false, 0, nil
	}
	if num,err := o.QueryTable("order_item").Filter("user_id",user.Id).Filter("buy_status",0).RelatedSel().All(&allItem); err != nil {
		beego.Error("get all buyuser error ",err)
		return false, num, nil
	}else {
		return true, num, allItem
	}
}

//查询购票信息
func QueryBuyUser(buyName string) (bool, int64, []OrderItem){
	o := orm.NewOrm()
	var allItem []OrderItem
	var num int64
	var err error
	if buyName == "" {
		if num,err = o.QueryTable("order_item").RelatedSel().All(&allItem); err != nil {
			return false, 0, nil
		}
	}else {
		if num,err = o.QueryTable("order_item").Filter("name",buyName).RelatedSel().All(&allItem); err != nil {
			return false, 0, nil
		}
	}
	return true, num, allItem
}






