package class

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type AirTicket struct {
	Id 			int
	Name 		string
	Time 		string
	Price 		float64
	Origin 		string 		`orm:"null"`//始发地
	Destination string 		`orm:"null"`//目的地
	Seats 		int
	Flight 		*Flights	`orm:"rel(fk)"`
}

//添加航班
func (c *AirTicket)AddAirTicket()  bool{
	o := orm.NewOrm()
	airName := AirTicket{Name : c.Name}
	if err := o.Read(&airName); err == nil {
		fmt.Println(0)
		return false
	}
	fmt.Println(1)
	var flight Flights
	if err := o.QueryTable("flights").Filter("name",c.Flight.Name).One(&flight); err != nil {
		fmt.Println(2)
		return false
	}else {
		c.Flight.Id = flight.Id
		fmt.Println(3)
	}
	fmt.Println(4)
	if _, err := o.Insert(c); err != nil {
		beego.Error("add airticket error ",err)
		return false
	}else {
		return true
	}
}

//删除航班
func DeleteAirTicket(name string) bool{
	o := orm.NewOrm()
	if _, err := o.QueryTable("air_ticket").Filter("name",name).Delete(); err != nil {
		beego.Error("delete airticket error ",err)
		return false
	}else {
		return true
	}
}

//修改航班
func (c *AirTicket)UpdateAirTicket() bool{
	o := orm.NewOrm()
	airTicket := AirTicket{Id : c.Id}
	if err := o.Read(&airTicket); err != nil {
		beego.Error("updateAirTicket get air error ",err)
		return false
	}else {
		if _, err := o.Update(c); err != nil {
			beego.Error("updateAriTicket error ",err)
			return false
		}
	}
	return true
}

//获取航班信息
func GetAllAirTicket() (bool, int64, []AirTicket) {
	o := orm.NewOrm()
	var airTicket []AirTicket
	if num, err := o.QueryTable("air_ticket").RelatedSel().All(&airTicket); err != nil {
		return false,0,nil
	}else {
		return true,num,airTicket
	}
}

//查询航班
func QueryAirTicket(name string) (bool, int64, []AirTicket){
	o := orm.NewOrm()
	var airTicket []AirTicket
	var id Flights
	if  err := o.QueryTable("flights").Filter("name",name).One(&id); err != nil {
		return false,0,nil
	}
	if num,err := o.QueryTable("air_ticket").Filter("flight_id",id.Id).RelatedSel().All(&airTicket); err != nil {
		beego.Error("query airticket error ",err)
		return false, 0, nil
	}else {
		return true, num, airTicket
	}

}
