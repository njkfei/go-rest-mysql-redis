package models

import (
	//"database/sql"
	"github.com/astaxie/goredis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	 "encoding/json"
)
var (
	client goredis.Client
	user string
	password string
	host string
	db string
	dns string
	posts []Postinfo
)

func init() {
	user = beego.AppConfig.String("mysqluser")
	password = beego.AppConfig.String("mysqlpass")
	host = beego.AppConfig.String("mysqlurls")
	db = beego.AppConfig.String("mysqldb")
	dns = beego.AppConfig.String("mysqldns")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql",dns)


	// register model
	//orm.RegisterModel(new(Postinfo))
	orm.RegisterModel(new (User))

	// create table
	orm.RunSyncdb("default", false, true)

	client.Addr = "127.0.0.1:6379"

}

type User struct {
	Id int32 `orm:"pk"`
}

type Postinfo struct {
	Id       int32 `orm:"pk"`
	PackageName string
	Version string
	Title string
	DownloadUrl string
	PreviewImageUrl string
}


func GetAllPostinfos() []Postinfo {

	var (
		val []byte
	 	err error
		post Postinfo

		maps []orm.Params
	)
	val, err = client.Get("themes");

	if  err != nil || string(val) != "hello" {
		beego.Notice("get failed")
		o := orm.NewOrm()
		o.Using("yii2basic")


		num,err1 := o.Raw("SELECT `id`, `pacname` as `packageName` ,`version`,`title`,`zip_source` as `downloadUrl`,`theme_url` as `previewImageUrl`  FROM `postinfo` where `status`=1").QueryRows(&posts)

		if err1 != nil {
			beego.Notice("select error")
		}
		beego.Notice(num)
		beego.Notice(posts)
		num, err := o.Raw("SELECT  `pacname` as `packageName` ,`version`,`title`,`zip_source` as `downloadUrl`,`theme_url` as `previewImageUrl`  FROM `postinfo` where `status`=1").Values(&maps)
		if err == nil && num > 0 {
			var index int64
			beego.Notice(num)
			for index = 0; index < num; index++ {
				beego.Notice(maps[index]) // slene
				post = maptoPostinfo(maps[index])
				posts[index] = post
				//FillStruct(maps[index],post[index])
				//beego.Notice(post[index])
			}
		}

		beego.Notice(maps)
		json,err := json.Marshal(posts);
		//beego.Notice(json)
		if err == nil {
			err = client.Set("themes", json)
		}

		return posts
	}


	err = json.Unmarshal(val,posts)

	return posts
}

func GetPostinfo(id string)(ret Postinfo)  {
	var post Postinfo
	o := orm.NewOrm()
	o.Using("yii2basic")

	 err := o.Raw("SELECT `id`, `pacname` as `packageName` ,`version`,`title`,`zip_source` as `downloadUrl`,`theme_url` as `previewImageUrl`  FROM `postinfo` where `status`=1 and id = ?", id).QueryRow(&post)

	if err == nil {
		beego.Notice("select error")
	}

	beego.Notice(post)

	return post
}

func ClearRedis()[]Postinfo{
	_,_ = client.Del("themes");


	return GetAllPostinfos()
}

func maptoPostinfo(maps map[string] interface{})(post Postinfo){
	if str,ok := maps[`packageName`].(string); ok{
		post.PackageName = str
	}

	if str,ok := maps[`downloadUrl`].(string); ok{
		post.DownloadUrl = str
	}

	if str,ok := maps[`previewImageUrl`].(string); ok{
		post.PreviewImageUrl = str
	}

	if str,ok := maps[`title`].(string); ok{
		post.Title = str
	}

	if str,ok := maps[`version`].(string); ok{
		post.Version = str
	}

	return  post
}

func GetCaches()( []Postinfo)  {
	var (
	//	post Postinfo
		maps []orm.Params
	)

	if(len(posts) > 0){
		return posts
	}

	o := orm.NewOrm()
	o.Using("yii2basic")
	num, err := o.Raw("SELECT  `pacname` as `packageName` ,`version`,`title`,`zip_source` as `downloadUrl`,`theme_url` as `previewImageUrl`  FROM `postinfo` where `status`=1").Values(&maps)
	if err == nil && num > 0 {
		var index int64
		beego.Notice(num)
		for index = 0; index < num; index++ {
			beego.Notice(maps[index]) // slene
			post := maptoPostinfo(maps[index])
			posts[index] = post
		}
	}

	return posts
}