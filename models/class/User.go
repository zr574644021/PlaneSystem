package class

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Users struct {
	Id 		 int
	UserName string `orm:"unique"`
	PassWord string
	Other 	 string
	State    int	`orm:"default(0)"`
}

type Admin struct {
	Id			int
	UserName 	string `orm:"unique"`
	PassWord	string
}

//旅客登录
func (c *Users)UserLoign()  bool{
	o := orm.NewOrm()
	var user Users
	if err := o.QueryTable("users").Filter("user_name",c.UserName).One(&user); err != nil {
		beego.Error("login error ", err)
		return false
	}
	if c.PassWord == user.PassWord {
		return true
	}else {
		return false
	}
}

//管理员登录
func AdminLogin(admin Admin) bool{
	o := orm.NewOrm()
	var user Admin
	if err := o.QueryTable("admin").Filter("user_name",admin.UserName).One(&user); err != nil {
		beego.Error("login error ", err)
		return false
	}
	if admin.PassWord == user.PassWord {
		return true
	}else {
		return false
	}
}

//添加用户
func (c *Users)AddUsers() bool{
	o := orm.NewOrm()
	var users Users
	if err := o.QueryTable("users").Filter("user_name",c.UserName).One(&users); err == nil {
		if users.State == 1 {
			users.State = 0
			users.PassWord = c.PassWord
			users.Other = c.Other
			//users.Money = c.Money
			if _,err := o.Update(&users); err != nil {
				beego.Error("insert update user error ",err)
				return false
			}
			return true
		}
		return false
	}else {
		if _, err := o.Insert(c);err != nil {
			beego.Error("insert user error ",err)
			return false
		}
	}
	return true
}


//获取用户信息
func GetUsers(identify int) (int64, []Users, error){
	var userAll []Users
	o := orm.NewOrm()
	if i, err := o.QueryTable("users").Filter("identify",identify).All(&userAll); err != nil {
		beego.Error("get user error ",err)
		return 0, nil, err
	}else {
		return i, userAll, nil
	}
}

//删除用户信息(软删除)
func DeleteUsers(username string) bool{
	o := orm.NewOrm()
	var user Users
	if err := o.QueryTable("users").Filter("username",username).Filter("state",0).One(&user); err != nil {
		beego.Error("deleteuser get user error ", err)
		return false
	}else {
		user.State = 1
		if _,err = o.Update(&user); err != nil {
			beego.Error("deleteuser user error ", err)
			return false
		}
	}
	return true
}

//修改用户信息
func (c *Users)UpdateUser() bool{
	o := orm.NewOrm()
	user := Users{UserName : c.UserName}
	if err := o.Read(&user); err != nil {
		beego.Error("updateUser get user error ", err)
		return false
	}else {
		c.State = 0
		if _, err := o.Update(&c); err != nil {
			beego.Error("updateUser user error ", err)
			return false
		}
	}
	return true
}


//查询用户信息
func QueryUser(username string) (bool, Users){
	o := orm.NewOrm()
	var user Users
	if err := o.QueryTable("users").Filter("username",username).One(&user); err != nil {
		beego.Error("query user error ",err)
		return false, user
	}else {
		return true, user
	}
}