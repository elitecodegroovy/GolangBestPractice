package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "https://github.com/elitecodegroovy"
	c.Data["Info"] = "Welcome Golang World"
	c.Data["Email"] = "org.jiang@gmail.com"
	c.TplName = "index.tpl"
}
