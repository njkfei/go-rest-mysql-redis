package models

import (
	//"database/sql"
	"github.com/astaxie/goredis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)
var (
	PostinfoList map[string]*Postinfo
	client goredis.Client
)

func init() {
	PostinfoList = make(map[string]*Postinfo)
	u := Postinfo{"1", "astaxie", "11111",1,"test","test","test","test","test",2}
	PostinfoList["1"] = &u

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:rootroot@/yii2basic?charset=utf8")


	// register model
	orm.RegisterModel(new(Postinfo))

	// create table
	orm.RunSyncdb("default", false, true)

	client.Addr = "127.0.0.1:6379"

}

/*
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `pacname` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `version` varchar(30) DEFAULT NULL,
  `version_in` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL,
  `zip_source` varchar(200) DEFAULT NULL,
  `zip_name` varchar(200) DEFAULT NULL,
  `themepic` varchar(200) DEFAULT NULL,
  `theme_url` varchar(100) DEFAULT NULL,
  `status` int(5) DEFAULT '0',
  PRIMARY KEY (`id`)
 */

type Postinfo struct {
	Id       string `orm:"pk"`
	Pacname string
	Version string
	Version_in int32
	Title string
	Zip_source string
	Zip_name string
	Theme_pic string
	Theme_url string
	Status int32
}


func GetAllPostinfos() []Postinfo {
/*	db, err := sql.Open("mysql", "root:rootroot@/yii2basic")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("SELECT * FROM postinfo") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	beego.Notice(stmtIns)*/

	var val []byte
	var err error

	err = client.Set("a", []byte("hello"))

	if err != nil {
		beego.Notice("set failed", err.Error())
	}

	if val, err = client.Get("a"); err != nil || string(val) != "hello" {
		beego.Notice("get failed")
	}

	beego.Notice(val)

	if typ, err := client.Type("a"); err != nil || typ != "string" {
		beego.Notice("type failed", typ)
	}

	//if keys, err := client.Keys("*"); err != nil || len(keys) != 1 {
	//    t.Notice("keys failed", keys)
	//}

	//client.Del("a")

	if ok, _ := client.Exists("a"); ok {
		beego.Notice("Should be deleted")
	}

	var post []Postinfo
	o := orm.NewOrm()
	o.Using("yii2basic")

	num,err1 := o.Raw("SELECT * FROM postinfo WHERE 1=1").QueryRows(&post)

	if err1 == nil {
		beego.Notice("select error")
	}
	beego.Notice(num)

	beego.Notice(post)

	return post
}

func GetPostinfo(id string)(ret Postinfo)  {
	var post Postinfo
	o := orm.NewOrm()
	o.Using("yii2basic")

	err := o.Raw("SELECT * FROM postinfo WHERE id = ?", id).QueryRow(&post)

	if err == nil {
		beego.Notice("select error")
	}

	beego.Notice(post)

	return post
}
