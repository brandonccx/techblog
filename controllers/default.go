package controllers

import (
	"techblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Item struct {
	Title string
	Url   string
	Id    string
}

var menus []Item = []Item{
	Item{Title: "Home", Url: "/", Id: "home"},
	Item{Title: "Category", Url: "/category", Id: "category"},
	Item{Title: "Article", Url: "/article", Id: "article"},
}

func (c *MainController) Get() {
	var err error
	c.TplName = "index.html"
	c.Data["Menu"] = getMenu("home")
	c.Data["Title"] = "Tech Blog"
	c.Data["Menus"] = menus
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["Articles"], err = models.GetAllArticles("time")
	if err != nil {
		beego.Error(err)
	}
	beego.Emergency("get default controller")
}

func getMenu(id string) Item {
	for _, item := range menus {
		if item.Id == id {
			return item
		}
	}
	return Item{}
}
