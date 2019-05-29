package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 对应标签的结构体
type Label struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
type Art struct {
	LabelId    int    `json:"label_id"`
	ArticleId  int    `json:"article_id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type LabelPage struct {
	LabelList []Label
	Cnt       int64
}

func init() {
	path := strings.Join([]string{"root", ":", "root123", "@tcp(", "62.234.14.42", ":", "3306", ")/", "wt_app", "?charset=utf8"}, "")
	fmt.Println(path)
	orm.RegisterDataBase("default", "mysql", path)

	orm.RegisterModelWithPrefix("wt_", new(Label))

	fmt.Println("数据库初始化成功...")
}

func main() {
	o := orm.NewOrm()
	// o.Using("default")
	// query := o.QueryTable("wt_label")
	// var labelPage LabelPage
	// query.Limit(1).All(&labelPage.LabelList)
	// labelPage.Cnt, _ = query.Count()
	// fmt.Printf("结果： %+v", labelPage)
	type labelsArt struct {
		Label
		Art
		CreateTimeOne string
		CreateTimeTwo string
	}
	// var label labelsArt
	// res := make(orm.Params)
	var labels []labelsArt
	res, err := o.Raw("select a.*, a.create_time as create_time_one, b.*  ,b.create_time as create_time_two from wt_label a left join wt_art_label b on a.id = b.label_id where a.id = ?", 2).QueryRows(&labels)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	fmt.Printf("%#v", labels)
}
