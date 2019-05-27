package controllers

import (
	"strings"
	"time"
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

//GetAll 修改文章
func (a *GoodsController) UpdateArticle() {
	id, _ := a.GetInt("id")
	if id == 0 {
		goods := new(models.Goods)
		goods.Name = strings.TrimSpace(a.GetString("name"))
		goods.Abstract = a.GetString("abstract")
		goods.Type, _ = a.GetInt("type")
		goods.OriginPrice, _ = a.GetFloat("origin_price")
		goods.CurrentPrice, _ = a.GetFloat("current_price")
		goods.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		goods.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
		if _, err := goods.CreateOne(); err != nil {
			a.Data["json"] = common.Fail(300, "创建文章失败")
		}
		a.Data["json"] = common.Fail(200, "创建文章成功")
	} else {
		goods := models.GetGoodsOne(id)
		goods.Name = strings.TrimSpace(a.GetString("name"))
		goods.Abstract = a.GetString("abstract")
		goods.Type, _ = a.GetInt("type")
		goods.OriginPrice, _ = a.GetFloat("origin_price")
		goods.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
		if _, err := goods.UpdateOne(); err != nil {
			a.Data["json"] = common.Fail(300, "修改文章失败")
		}
		a.Data["json"] = common.Fail(200, "修改文章成功")
	}
	a.ServeJSON()

}
