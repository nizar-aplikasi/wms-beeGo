package controllers

import (
	"github.com/astaxie/beego"
	Dashboard "wms/repo/dashboard"
	"fmt"
)


type MainController struct {
	beego.Controller
}

func (c *MainController) Get(){
	v := c.GetSession("controllerSession")
    if v == nil {
		fmt.Println("session kosong")
		c.Redirect("/", 302)
	}
	fmt.Println("session ada")
	repo := Dashboard.RepoDashboard{}

	totalIn, errTotalIn := repo.GetTotalIncomingGoods()

	if errTotalIn != nil {
		c.Data["errormsg"] = errTotalIn
		c.Layout = "template.html"
		c.TplName = "dashboard.html" //buat load halaman
	}

	totalOut, errTotalOut := repo.GetTotalOutcomingGoods()

	if errTotalOut != nil {
		c.Data["errormsg"] = errTotalOut
		c.Layout = "template.html"
		c.TplName = "dashboard.html" //buat load halaman
	}

	totalUser, errTotalUser := repo.GetTotalUser()

	if errTotalUser != nil {
		c.Data["errormsg"] = errTotalUser
		c.Layout = "template.html"
		c.TplName = "dashboard.html" //buat load halaman
	}

	result, errMax := repo.Max()

	if errMax != "	" {
		c.Data["errormsg"] = errMax
		c.Layout = "template.html"
		c.TplName = "dashboard.html" //buat load halaman
	}

	c.Data["jumlahMax"] = result
	c.Data["jumlahMasuk"] = totalIn
	c.Data["jumlahKeluar"] = totalOut
	c.Data["jumlahUser"] = totalUser
	c.Layout = "template.html"
	c.TplName = "dashboard.html" //buat load halaman
}


