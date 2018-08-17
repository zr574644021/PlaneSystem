package controllers

import (
	"airTicket/models/class"
	"encoding/json"
	"fmt"
)

type AirController struct {
	class.AirTicket
	MainController
}

type queryName struct {
	Name string
}

//@router /api/airQuery [*]
func (c *AirController) AirLoad() {
	if LoginCheck(&c.MainController) {
		var flightName queryName
		json.Unmarshal(c.Ctx.Input.RequestBody,&flightName)
		if flightName.Name == "" {
			flag, num, airTicket := class.GetAllAirTicket()
			if flag {
				result := ResultJson(4000)
				result["record"] = airTicket
				result["num"] = num
				c.Data["json"] = result
			}else {
				c.Data["json"] = ResultJson(4004)
			}
		}else {
			flag, num, airTicket := class.QueryAirTicket(flightName.Name)
			if flag {
				result := ResultJson(4000)
				result["record"] = airTicket
				result["num"] = num
				c.Data["json"] = result
			}else {
				c.Data["json"] = ResultJson(4004)
			}
		}
		c.ServeJSON()
		return
	}
}

//@router /api/airAdd [*]
func (c *AirController)AirAdd () {
	if LoginCheck(&c.MainController) {
		json.Unmarshal(c.Ctx.Input.RequestBody,&c)
		fmt.Println(c.Flight.Name)
		if  c.AddAirTicket(){
			c.Data["json"] = ResultJson(4000)
		}else {
			c.Data["json"] = ResultJson(4004)
		}
		c.ServeJSON()
		return
	}
}

//@router /api/airUpdate [*]
func (c *AirController)AirUpdate () {
	if LoginCheck(&c.MainController){
		json.Unmarshal(c.Ctx.Input.RequestBody,&c)
		if c.UpdateAirTicket() {
			c.Data["json"] = ResultJson(4000)
		}else {
			c.Data["json"] = ResultJson(4004)
		}
		c.ServeJSON()
		return
	}
}

//@router /api/airDelete [*]
func (c *AirController) AirDelete() {
	if LoginCheck(&c.MainController) {
		json.Unmarshal(c.Ctx.Input.RequestBody,&c)
		if class.DeleteAirTicket(c.Name) {
			c.Data["json"] = ResultJson(4000)
		}else {
			c.Data["json"] = ResultJson(4004)
		}
		c.ServeJSON()
		return
	}
}

