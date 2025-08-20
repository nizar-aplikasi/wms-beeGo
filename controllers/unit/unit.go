package controllers

import (
	"github.com/astaxie/beego"
	Units "wms/repo/units"
	"wms/models"
	"strconv"
	"fmt"
)

type UnitController struct {
	beego.Controller
}

func (c *UnitController) All() {
	//Check Session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }	
	repo := Units.UnitRepo{}
	result, err := repo.GetAll()

	if err != nil {
		c.Data["units"] = err
		c.Layout = "template.html"
		c.TplName = "unit.html"
	}

	c.Data["units"] = result
	c.Layout = "template.html"
	c.TplName = "unit.html"
	
}

func (c *UnitController) Add() {	
	repo := Units.UnitRepo{}
	unitModel := models.Units{}

	unitModel.Code = c.GetString("kode")
	unitModel.Name = c.GetString("nama")

	result, err := repo.Add(unitModel)

	if err != nil {
		c.Data["errormsg"] = err
		c.Layout = "template.html"
		c.TplName = "form-unit.html"
	}
	c.Data["errormsg"] = result
	c.Redirect("/units", 302)
}

func (c *UnitController) Detail() {	
	//Check Session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	repo := Units.UnitRepo{}
	id := c.Ctx.Input.Param(":id")
	i, _ := strconv.Atoi(id)
	result, err := repo.GetById(i)

	if err != nil {
		c.Data["errormsg"] = err
		c.Layout = "template.html"
		c.TplName = "form-unit.html"
	}
	c.Data["units"] = result
	c.Layout = "template.html"
	c.TplName = "form-unit.html"
}

func (c *UnitController) Update() {
	unitModel := models.Units{}
	unitRepo := Units.UnitRepo{}

	id := c.Ctx.Input.Param(":id")
	i, _ := strconv.Atoi(id)
	unitModel.Id = i

	unitModel.Code 	= c.GetString("kode")
	unitModel.Name 	= c.GetString("nama")
	
	fmt.Println(unitModel)
	result, err := unitRepo.Update(unitModel)
	if err != nil {
		c.Data["errormsg"] = err
		c.Redirect("/units",302)
	}
	c.Data["errormsg"] = result
	c.Redirect("/units",302)
}

func (c *UnitController) Delete() {
	//Check Session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	unitRepo := Units.UnitRepo{}

	id := c.Ctx.Input.Param(":id")
	i, _ := strconv.Atoi(id)

	result, err := unitRepo.Delete(i)

	if err != nil {
		c.Data["errormsg"] = err
		c.Redirect("/units",302)
	}
	c.Data["pesan"] = result
	c.Redirect("/units",302)
}