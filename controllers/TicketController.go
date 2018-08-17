package controllers

import (
	"airTicket/models/class"
	"encoding/json"
	"fmt"
)

type TicketController struct {
	MainController
}


//@router /api/ticketQuery [*]
func (c *TicketController)QueryTicket() {
	if LoginCheck(&c.MainController) {
		var buyName buyName
		json.Unmarshal(c.Ctx.Input.RequestBody,&buyName)

		if buyName.Name == "" {
			flag, num, orderItem := class.AdminGetAllBuyUser()
			if  flag {
				result := ResultJson(4000)
				result["record"] = orderItem
				result["number"] = num
				c.Data["json"] = result
			}else {
				c.Data["json"] = ResultJson(4004)
			}
		}else {
			flag, num, orderItem := class.QueryBuyUser(buyName.Name)
			if flag {
				result := ResultJson(4000)
				result["record"] = orderItem
				result["number"] = num
				c.Data["json"] = result
			}else {
				c.Data["json"] = ResultJson(4004)
			}
		}
		c.ServeJSON()
		return
	}
}

//@router /api/ticketDY [*]
func (c *TicketController)DYTicket() {
	if LoginCheck(&c.MainController) {
		var dyItem class.OrderItem
		json.Unmarshal(c.Ctx.Input.RequestBody,&dyItem)
		fmt.Println(dyItem)
		if class.DYResult(dyItem) {
			c.Data["json"] = ResultJson(4000)
		}else {
			c.Data["json"] = ResultJson(4004)
		}
		c.ServeJSON()
		return
	}
}

