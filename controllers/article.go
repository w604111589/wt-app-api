package controllers

import (
	"wt-app-api/common"
	"wt-app-api/models"

	"github.com/astaxie/beego"
)

//ArticleController  fdsf
type ArticleController struct {
	beego.Controller
}

//Get 获取文章
func (a *ArticleController) Get() {
	id, _ := a.GetInt("id", 1)
	article := models.GetArticleOne(id)
	res := common.Success(article)
	a.Data["json"] = res
	a.ServeJSON()
}

//GetAll 获取文章
func (a *ArticleController) GetAll() {
	page, _ := a.GetInt("page", 1)
	limit, _ := a.GetInt("limit", 10)
	title := a.GetString("title")
	abstract := a.GetString("abstract")
	status, _ := a.GetInt("status")
	filters := make(map[string]interface{})

	if title != "" {
		filters["title__contains"] = title
	}

	if abstract != "" {
		filters["abstract__contains"] = abstract
	}

	if status != 0 {
		filters["status"] = status
	}

	articlePage := models.GetArticleAll(page, limit, filters)

	res := common.Success(articlePage)
	a.Data["json"] = res
	a.ServeJSON()
}
