package controllers

type User struct {
	username string
	password string
}


func ResultJson(errcode int)  (result map[string]interface{}){
	result = make(map[string]interface{})
	result["status"] = errcode
	return result
}

func LoginCheck(c *MainController) bool{
	username := c.GetSession("username")
	password := c.GetSession("password")
	if username == nil || password == nil {
		c.Data["json"] = ResultJson(4001)
		c.ServeJSON()
		return false
	}else {
		return true
	}
}
