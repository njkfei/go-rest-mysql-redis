package controllers

import (
	"github.com/astaxie/beego"
	"kkapi/models"
)
// Operations about Users
type PostinfosController struct {
	beego.Controller
}


func (p PostinfosController) Get(){
	posts := models.GetAllPostinfos()
	p.Data["json"] = posts
	beego.Notice(posts)
	beego.Notice(p)
	p.ServeJSON()
}