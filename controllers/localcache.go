package controllers

import (
	"github.com/astaxie/beego"
	"kkapi/models"
)
// Operations about Users
type LocalController struct {
	beego.Controller
}

func (p LocalController) Get(){

	posts := models.GetCaches()
	p.Data["json"] = posts
	beego.Notice(posts)
	beego.Notice(p)
	p.ServeJSON()
}