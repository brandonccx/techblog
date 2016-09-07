package controllers

import (
	"techblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	var err error
	c.TplName = "category.html"
	c.Data["Title"] = "Tech Blog"
	c.Data["Menu"] = getMenu("category")
	c.Data["Menus"] = menus
	c.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}
}

func (c *CategoryController) Post() {
	var err error
	switch c.Ctx.Input.Param(":action") {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err = models.AddCategory(name)
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err = models.DelCategory(id)
		c.ServeJSON()
	}
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/category", 302)
	return
}
