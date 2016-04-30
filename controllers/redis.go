package controllers

import (
	"github.com/astaxie/beego"
	"kkapi/models"
)
// Operations about Users
type RedisController struct {
	beego.Controller
}


func (p RedisController) Get(){
	posts := models.ClearRedis()
	p.Data["json"] = posts
	beego.Notice(posts)
	beego.Notice(p)
	p.ServeJSON()
}

