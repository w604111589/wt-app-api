package models

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 初始化数据库
func init() {
	mysqlUser := beego.AppConfig.String("mysql_user")
	mysqlPass := beego.AppConfig.String("mysql_pass")
	mysqlDB := beego.AppConfig.String("mysql_db")
	mysqlHost := beego.AppConfig.String("mysql_host")
	mysqlPort := beego.AppConfig.String("mysql_port")

	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{mysqlUser, ":", mysqlPass, "@tcp(", mysqlHost, ":", mysqlPort, ")/", mysqlDB, "?charset=utf8"}, "")
	// path := strings.Join([]string{"root", ":", "root123", "@tcp(", "62.234.14.42", ":", "3306", ")/", "wt_app", "?charset=utf8"}, "")
	fmt.Println(path)
	orm.RegisterDataBase("default", "mysql", path)

	orm.RegisterModelWithPrefix("wt_", new(Label), new(Article), new(Goods), new(Comment))

	fmt.Println("数据库初始化成功...")
}
