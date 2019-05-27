package controllers

import (
	"fmt"
	"log"
	"time"
	"wt-app-api/common"
	"wt-app-api/models"

	"github.com/astaxie/beego"
)

//LabelController fsdfsdf/
type LabelController struct {
	beego.Controller
}

//Get fsdfsdf/
func (l *LabelController) Get() {
	id, _ := l.GetInt("id")

	label := models.GetOne(id)
	res := common.Success(label)

	l.Data["json"] = res
	l.ServeJSON()

}

func (l *LabelController) Test() {
	go func() {
		beego.Info(`这是一个错误 `)
		beego.Error(`这是一个错误111 `)
		log.Println("测试log")
	}()
	fmt.Println("当前时间：", beego.DateFormat(time.Now(), "2006-01-02 15:03:04"))
}

//GetAll fsdfsdf/
func (l *LabelController) GetAll() {
	page, _ := l.GetInt("page", 1)
	limit, _ := l.GetInt("limit", 10)
	name := l.GetString("name")
	status, err := l.GetInt("status")

	filters := make(map[string]interface{})

	if name != "" {
		filters["name__contains"] = name
	}

	if err == nil {
		filters["status"] = status
	}

	labelPage := models.GetAll(page, limit, filters)
	res := common.Success(labelPage)

	l.Data["json"] = res
	l.ServeJSON()

}
