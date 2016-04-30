package controllers

import (
	"github.com/astaxie/beego"
	"kkapi/models"
)
// Operations about Users
type PostinfoController struct {
	beego.Controller
}

func (p PostinfoController) Get(){
	id := p.Ctx.Input.Param(":id")

	posts := models.GetPostinfo(id)
	p.Data["json"] = posts
	beego.Notice(posts)
	beego.Notice(p)
	p.ServeJSON()
}