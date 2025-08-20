package controllers

import(
	"wms/models"
	Auth "wms/repo/auth"
	"github.com/astaxie/beego"
)

type MainController struct{
	beego.Controller
}


func(c *MainController) Login(){
	// flash := beego.ReadFromRequest(&c.Controller)
	flash := beego.NewFlash()
	repo := Auth.AuthRepo{}
	form := models.LoginForm{}
	form.Email = c.GetString("email")
	form.Password = c.GetString("password")
	result, text := repo.Login(form)
	if text == "Login Gagal"  {
		flash.Notice("Login gagal!")
		flash.Store(&c.Controller)
		c.Data["pesanError"] = "ada error"
		c.Redirect("/", 302)
		return
	} else {
		c.Data["dataUser"] = result
		c.Layout = "template.html"
		c.Redirect("/dashboard", 302)
	}
}




