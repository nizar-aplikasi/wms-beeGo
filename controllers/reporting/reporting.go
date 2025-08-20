package api

import (
	"github.com/astaxie/beego"
	Reporting "wms/repo/reporting"
	"fmt"
	WHelper"wms/helper"
)


type ReportingController struct {
	beego.Controller
}

func (c *ReportingController) All() {
	//cek session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	//flash
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["success"]; ok {
        // Display settings successful
        c.TplName = "reporting/success.html"
    } else if n, ok = flash.Data["error"]; ok {
        // Display error messages
        c.TplName = "reporting/success.html"
    } else {
		// Display default settings page
		fmt.Println(n, "Ini dari N")
        c.TplName = "reporting/success.html"
    }
	//logic
	repo := Reporting.ReportRepo{}
	result, err := repo.All()
	if err != nil {
		c.Data["errormsg"] = err
		c.Layout = "template.html"
		c.TplName = "reporting/in-out.html" //buat load halaman
	}
	fmt.Println(result)
	c.Data["list"] = result
	c.Layout = "template.html"
	c.TplName = "reporting/in-out.html" //buat load halaman
}

func (c *ReportingController) Export() {
	//Check Session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	flash := beego.NewFlash()
	exp := WHelper.HelperModule{}
	result := exp.ExportToExcell()
	fmt.Println(result)
	if(!result){
		flash.Error("Failed to generated excell file")
	}
	// c.Data["excellSuccess"] = result
	// c.Layout = "template.html"
	// c.TplName = "reporting/success-export.html" //buat load halaman
	flash.Success("Successfully Generate Excell File")
	// flash.Success("Success..")
	flash.Store(&c.Controller)
	c.Redirect("/reporting", 302)
}



