package models

import "github.com/astaxie/beego/orm"

// 对应标签的结构体
type Label struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type LabelPage struct {
	LabelList []*Label
	Cnt       int64
}

func GetOne(id int) Label {
	o := orm.NewOrm()
	o.Using("default")
	query := o.QueryTable("wt_label")
	var label Label
	query.Filter("id", id).One(&label)
	return label

}

func GetAll(page, limit int, filters map[string]interface{}) LabelPage {
	o := orm.NewOrm()
	o.Using("default")
	query := o.QueryTable("wt_label")
	if len(filters) > 0 {
		for k, v := range filters {
			query = query.Filter(k, v)
		}
	}
	var labelPage LabelPage
	query.Limit(limit, (page-1)*limit).All(&labelPage.LabelList)
	labelPage.Cnt, _ = query.Count()

	return labelPage
}
