package initial

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func InitMysql() {
	//从默认配置文件中取得参数，其中可以有某些默认缺省的，返回错误，然后在程序里面设置
	mysql_user := beego.AppConfig.String("mysqluser")
	mysql_password := beego.AppConfig.String("mysqlpass")
	mysql_url := beego.AppConfig.String("mysqlurl")
	mysql_port, err := beego.AppConfig.Int("mysqlport")
	mysql_database := beego.AppConfig.String("mysqldb")
	if nil != err {
		mysql_port = 3306
	}
	//注册数据库驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册数据库
	orm.RegisterDataBase("autotestopms", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", mysql_user, mysql_password, mysql_url, mysql_port, mysql_database))

}
