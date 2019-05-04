package controllers

import (
	"wt-app-api/common"
	"wt-app-api/models"

	"github.com/astaxie/beego"
)

//ArticleController  fdsf
type GoodsController struct {
	beego.Controller
}

//Get 获取文章
func (a *GoodsController) Get() {
	id, _ := a.GetInt("id", 1)
	goods := models.GetGoodsOne(id)
	res := common.Success(goods)
	a.Data["json"] = res
	a.ServeJSON()
}

//GetAll 获取文章
func (a *GoodsController) GetAll() {
	page, _ := a.GetInt("page", 1)
	limit, _ := a.GetInt("limit", 10)
	name := a.GetString("name")
	origin_price_gte, err_origin_gte := a.GetFloat("origin_price_gte")
	origin_price_lte, err_origin_lte := a.GetFloat("origin_price_lte")
	current_price_gte, err_current_gte := a.GetFloat("current_price_gte")
	current_price_lte, err_current_lte := a.GetFloat("current_price_lte")
	abstract := a.GetString("abstract")
	status, err := a.GetInt("status")
	filters := make(map[string]interface{})

	if err_origin_gte == nil && err_origin_lte == nil {
		filters["origin_price__gte"] = origin_price_gte
		filters["origin_price__lte"] = origin_price_lte
	}

	if err_current_gte == nil && err_current_lte == nil {
		filters["origin_current__gte"] = current_price_gte
		filters["origin_current__lte"] = current_price_lte
	}

	if name != "" {
		filters["name__contains"] = name
	}

	if abstract != "" {
		filters["abstract__contains"] = abstract
	}

	if err == nil {
		filters["status"] = status
	}

	goodsPage := models.GetGoodsAll(page, limit, filters)

	res := common.Success(goodsPage)
	a.Data["json"] = res
	a.ServeJSON()
}
