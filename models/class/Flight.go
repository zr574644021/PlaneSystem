package class

import "github.com/astaxie/beego/orm"

type Flights struct {
	Id int
	Name string `orm:"unique"`
}

func GetFlight()  (bool,[]Flights){
	o := orm.NewOrm()
	var allFlight []Flights
	if  _,err := o.QueryTable("flights").All(&allFlight); err != nil {
		return false, nil
	}else {
		return true, allFlight
	}
}