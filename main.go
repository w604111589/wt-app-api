package main

import (
	_ "wt-app-api/routers"

	_ "wt-app-api/models"

	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	os.MkdirAll("./logs", os.ModePerm)
	beego.SetLogger("file", `{"filename":"./logs/test.log"}`)
	beego.Run()
}
