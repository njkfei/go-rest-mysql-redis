# go-rest-mysql-redis
  go mysql redis restful api集成项目。
  
## 集成redis
```
	"github.com/astaxie/goredis"
```

## 集成mysql
```
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
```

## redis使用
```
var client goredis.Client
client.Addr = "127.0.0.1:6379"
val, err = client.Get("themes");
err = client.Set("themes", json)
```

## mysql使用
```
	dns = beego.AppConfig.String("mysqldns")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",dns)
	// register model
	//orm.RegisterModel(new(Postinfo))
	orm.RegisterModel(new (User))

		o := orm.NewOrm()
		o.Using("yii2basic")


		num,err1 := o.Raw("SELECT `id`, `pacname` as `packageName` ,`version`,`title`,`zip_source` as `downloadUrl`,`theme_url` as `previewImageUrl`  FROM `postinfo` where `status`=1").QueryRows(&posts)

		if err1 != nil {
			beego.Notice("select error")
		}
		beego.Notice(num)
		beego.Notice(posts)
		num, err := o.Raw("SELECT  `pacname` as `packageName` ,`version`,`title`,`zip_source` as `downloadUrl`,`theme_url` as `previewImageUrl`  FROM `postinfo` where `status`=1").Values(&maps)
```


## 测试结论
* 内存比redis快一个数量级
* redis比mysql快一个数量级
