package controllers

import (
	"github.com/astaxie/beego"
	"airTicket/models/class"
	"encoding/json"
)

type MainController struct {
	class.Users
	beego.Controller
}


/*func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}*/

type Login struct {
	Username string
	Password string
}


//@router /userLogin [*]
func (c *MainController) UsLogin() {
	var login Login
	json.Unmarshal(c.Ctx.Input.RequestBody, &login)
	c.UserName = login.Username
	c.PassWord = login.Password

	if c.UserLoign() {
		c.SetSession("username",c.UserName)
		c.SetSession("password",c.PassWord)
		c.Data["json"] = ResultJson(4000)//登录成功
	}else {
		c.Data["json"] = ResultJson(4004)//登录失败
	}
	c.ServeJSON()
	return
}


//@router /login [*]
func (c *MainController) AdLogin() {
	var login class.Admin
	json.Unmarshal(c.Ctx.Input.RequestBody, &login)
	if class.AdminLogin(login) {
		c.SetSession("username",c.UserName)
		c.SetSession("password",c.PassWord)
		c.Data["json"] = ResultJson(4000)//登录成功
	}else {
		c.Data["json"] = ResultJson(4004)//登录失败
	}
	c.ServeJSON()
	return
}

//@router /register [*]
func (c *MainController) Register()  {
	json.Unmarshal(c.Ctx.Input.RequestBody,&c)
	if c.AddUsers() {
		c.Data["json"] = ResultJson(4000)
	}else {
		c.Data["json"] = ResultJson(4004)
	}
	c.ServeJSON()
	return
}

//@router /updateUser [*]
func (c *MainController) Update() {
	c.UserName = c.GetString("username")
	c.PassWord = c.GetString("password")
	c.Other = c.GetString("other")

	if c.UpdateUser() {
		c.Data["json"] = ResultJson(4000)
	}else {
		c.Data["json"] = ResultJson(4004)
	}
	c.ServeJSON()
	return
}

//@router /api/getFlight [*]
func (c *MainController)GetAllFlight()  {
	flag, allFlight := class.GetFlight()
	if flag {
		result := ResultJson(4000)
		result["record"] = allFlight
		c.Data["json"] = result
	}else {
		c.Data["json"] = ResultJson(4004)
	}
	c.ServeJSON()
	return
}