package controllers

import (
	"wt-app-api/common"
	"wt-app-api/models"

	"github.com/astaxie/beego"
)

//ArticleController  fdsf
type CommentController struct {
	beego.Controller
}

//Get 获取文章评论
func (a *CommentController) Get() {
	id, _ := a.GetInt("id", 1)
	article := models.GetCommentOne(id)
	res := common.Success(article)
	a.Data["json"] = res
	a.ServeJSON()
}

//Create 获取文章评论
func (a *CommentController) Create() {
	id, _ := a.GetInt("id", 1)
	article := models.GetCommentOne(id)
	res := common.Success(article)
	a.Data["json"] = res
	a.ServeJSON()
}

//GetAll 获取文章评论
func (a *CommentController) GetAll() {
	page, _ := a.GetInt("page", 1)
	limit, _ := a.GetInt("limit", 10)
	article_id, err_ariticle_id := a.GetInt("article_id", 1)
	// title := a.GetString("title")
	content := a.GetString("content")
	status, err_status := a.GetInt("status")
	filters := make(map[string]interface{})

	if err_ariticle_id == nil {
		filters["article_id"] = article_id
	}

	if content != "" {
		filters["content__contains"] = content
	}

	if err_status == nil {
		filters["status"] = status
	}

	commentPage := models.GetCommentAll(page, limit, filters)

	res := common.Success(commentPage)
	a.Data["json"] = res
	a.ServeJSON()
}
