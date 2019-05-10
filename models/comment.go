package models

import (
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id            int    `json:"id" form:"-"`
	ArticleId     int    `json:"article_id" form:"article_id"`
	Content       string `json:"content" form:"content"`
	Commentor     string `json:"commentor" form:"commentor"`
	Comment_lever string `json:"commentor_lever" form:"commentor_lever"`
	CreateTime    string `json:"create_time"`
}

func GetCommentOne(id int) Comment {
	o := orm.NewOrm()
	o.Using("default")
	var comment Comment
	o.QueryTable("wt_article").Filter("id", id).One(&comment)
	return comment
}

func CreateCommentOne(comment Comment) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(comment)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// 分页时的文章结构体
type CommentPage struct {
	CommentList []*Comment
	Cnt         int64
}

// 分页获取文章参数
func GetCommentAll(page, limit int, filters map[string]interface{}) CommentPage {

	o := orm.NewOrm()
	o.Using("default")
	var commentPage CommentPage
	offset := (page - 1) * limit
	query := o.QueryTable("wt_comment")
	if len(filters) > 0 {
		for k, v := range filters {
			query = query.Filter(k, v)
		}
	}
	// query = query.Limit(limit, offset)
	commentPage.Cnt, _ = query.Count()
	query.Limit(limit, offset).All(&commentPage.CommentList)

	return commentPage
}
