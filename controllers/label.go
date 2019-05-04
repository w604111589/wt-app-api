package controllers

import (
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
