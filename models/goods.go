package models

import (
	"github.com/astaxie/beego/orm"
)

//Goods 商品的结构体
type Goods struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Type         int     `json:"type"`
	Abstract     string  `json:"abstract"`
	OriginPrice  float32 `json:"origin_price"`
	CurrentPrice float32 `json:"current_price"`
	ImageUrl     string  `json:"image_url"`
	CreateTime   string  `json:"create_time"`
	UpdateTime   string  `json:"update_time"`
}

//GetGoodsOne 通过获取单条记录
func GetGoodsOne(id int) Goods {
	o := orm.NewOrm()
	o.Using("default")
	var goods Goods
	o.QueryTable("wt_goods").Filter("id", id).One(&goods)
	return goods
}

//GoodsPage 分页时的文章结构体
type GoodsPage struct {
	Lists []*Goods `json:"lists"`
	Total int64    `json:"total"`
}

//GetGoodsAll 分页获取文章参数
func GetGoodsAll(page, limit int, filters map[string]interface{}) GoodsPage {

	o := orm.NewOrm()
	o.Using("default")
	var goodsPage GoodsPage
	offset := (page - 1) * limit
	query := o.QueryTable("wt_goods")
	if len(filters) > 0 {
		for k, v := range filters {
			query = query.Filter(k, v)
		}
	}
	// query = query.Limit(limit, offset)
	goodsPage.Total, _ = query.Count()
	query.Limit(limit, offset).All(&goodsPage.Lists)

	return goodsPage
}
