package models

import (
	"os"
	"path"
	"strconv"
	"time"

	util "github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME        = "data/techblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id                int64
	Title             string
	Views             int64     `orm:"index"`
	CreatedTime       time.Time `orm:"index"`
	ArticleTime       time.Time `orm:"null;index"`
	ArticleCount      int64
	ArticleLastUserId int64
}

type Article struct {
	Id                int64
	Uid               int64
	Title             string
	Content           string `orm:"size(5000)"`
	Attachment        string
	CreatedTime       time.Time `orm:"index"`
	UpdatedTime       time.Time `orm:"index"`
	Views             int64     `orm:"index"`
	Author            string
	RepliedTime       time.Time `orm:"null;index"`
	RepliedCount      int64
	RepliedLastUserId int64
}

func RegisterDB() {
	if !util.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Article))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddArticle(title, content string) error {
	var err error
	o := orm.NewOrm()
	article := &Article{
		Title:       title,
		Content:     content,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	_, err = o.Insert(article)
	if err != nil {
		return err
	}
	return nil
}

/**
 * Get all articles
 * @param order {String} order by ...
 */
func GetAllArticles(order string) ([]*Article, error) {
	var err error
	o := orm.NewOrm()
	articles := make([]*Article, 0)
	qs := o.QueryTable("article")

	if order == "time" {
		_, err = qs.OrderBy("-CreatedTime").All(&articles)
	} else {
		_, err = qs.All(&articles)
	}
	return articles, err
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{
		Title:       name,
		CreatedTime: time.Now(),
	}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)

	if err != nil {
		return err
	}

	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
