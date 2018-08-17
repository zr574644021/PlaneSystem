package controllers

import (
	"airTicket/models/class"
	"encoding/json"
	"fmt"
)


type CarController struct {
	MainController
}

type buyName struct {
	Name string
}


//@router /api/ResBuyCar [*]
func (c *CarController)BuyTicket() {
	if LoginCheck(&c.MainController) {
		var tableData []class.OrderItem
		json.Unmarshal(c.Ctx.Input.RequestBody,&tableData)
		fmt.Println(tableData)
		if class.BuyCarResult(tableData) {
			c.Data["json"] = ResultJson(4000)
		}else {
			c.Data["json"] = ResultJson(4004)
		}
		c.ServeJSON()
		return
	}
}

//@router /api/GetBuyCar [*]
func (c *CarController)GetAllCar()  {
	if LoginCheck(&c.MainController) {
		var buyName buyName
		json.Unmarshal(c.Ctx.Input.RequestBody,&buyName)
		fmt.Println(buyName)
		flag, _, carItem := class.GetCarMessage(buyName.Name)
		if  flag {
			result := ResultJson(4000)
			result["record"] = carItem
			c.Data["json"] = result
		}else {
			c.Data["json"] = ResultJson(4004)
		}
		c.ServeJSON()
		return
	}
}

//@router /api/GetOrder [*]
func (c *CarController)GetAllOrder()  {
	if LoginCheck(&c.MainController) {
		var buyName buyName
		json.Unmarshal(c.Ctx.Input.RequestBody,&buyName)
		fmt.Println(buyName)
		flag, _, carItem := class.GetOrder(buyName.Name)
		if  flag {
			result := ResultJson(4000)
			result["record"] = carItem
			c.Data["json"] = result
		}else {
			c.Data["json"] = ResultJson(4004)
		}
		c.ServeJSON()
		return
	}
}


//@router /api/AddBuyCar [*]
func (c *CarController)CarTicket()  {
	var buyUser class.OrderItem

	json.Unmarshal(c.Ctx.Input.RequestBody,&buyUser)
	if class.AddBuyCar(buyUser) {
		c.Data["json"] = ResultJson(4000)
	}else {
		c.Data["json"] = ResultJson(4004)
	}
	c.ServeJSON()
	return
}

//@router /api/DelBuyCar [*]
func (c *CarController)CarDel() {
	if LoginCheck(&c.MainController) {
		var del class.OrderItem
		json.Unmarshal(c.Ctx.Input.RequestBody,&del)
		fmt.Println(del)
		if class.DelBuyCar(del.Id,del.AirTic.Name) {
			c.Data["json"] = ResultJson(4000)
		}else {
			c.Data["json"] = ResultJson(4004)
		}
		c.ServeJSON()
		return
	}
}

