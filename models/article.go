package models

import (
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id              int    `json:"id"`
	CreaterId       string `json:"create_id"`
	Title           string `json:"title"`
	Abstract        string `json:"abstract"`
	Importance      string `json:"importance"`
	Type            string `json:"type"`
	Content         string `json:"content"`
	ReleaseTime     string `json:"release_time"`
	CommentDisabled string `json:"comment_disabled"`
	SourceUri       string `json:"source_uri"`
	ImageUri        string `json:"image_uri"`
	Platform        int    `json:"platform"`
	Status          int    `json:"status"`
	CreateTime      string `json:"create_time"`
	UpdateTime      string `json:"update_time"`
}

func GetArticleOne(id int) Article {
	o := orm.NewOrm()
	o.Using("default")
	var article Article
	o.QueryTable("wt_article").Filter("id", id).One(&article)
	return article
}

// 分页时的文章结构体
type ArticlePage struct {
	ArticleList []*Article
	Cnt         int64
}

// 分页获取文章参数
func GetArticleAll(page, limit int, filters map[string]interface{}) ArticlePage {

	o := orm.NewOrm()
	o.Using("default")
	var articlePage ArticlePage
	offset := (page - 1) * limit
	query := o.QueryTable("wt_article")
	if len(filters) > 0 {
		for k, v := range filters {
			query = query.Filter(k, v)
		}
	}
	// query = query.Limit(limit, offset)
	articlePage.Cnt, _ = query.Count()
	query.Limit(limit, offset).All(&articlePage.ArticleList)

	return articlePage
}
