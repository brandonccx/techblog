package controllers

import (
	"techblog/models"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	var err error
	c.TplName = "article.html"
	c.Data["Title"] = "Tech Blog"
	c.Data["Menu"] = getMenu("article")
	c.Data["Menus"] = menus
	c.Data["Articles"], err = models.GetAllArticles("id")
	if err != nil {
		beego.Error(err)
	}
}

func (c *ArticleController) Add() {
	c.TplName = "article_add.html"
	c.Data["Title"] = "Tech Blog"
	c.Data["Menu"] = getMenu("article")
	c.Data["Menus"] = menus
}

func (c *ArticleController) Post() {
	inputs := c.Input()
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	err := models.AddArticle(inputs.Get("title"), inputs.Get("content"))
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/article", 302)
}
