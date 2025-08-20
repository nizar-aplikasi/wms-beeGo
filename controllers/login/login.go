package controllers

import (
	"fmt"
	"wms/models"
	UserRepo "wms/repo/users"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

// Render halaman login
func (c *LoginController) Get() {
	c.TplName = "login.html"
}

// Proses logout
func (c *LoginController) Logout() {
	c.DelSession("controllerSession")
	c.Redirect("/", 302)
}

// Proses login
func (c *LoginController) Login() {
	repo := UserRepo.UsersRepo{}
	loginForm := models.LoginForm{
		Email:    c.GetString("email"),
		Password: c.GetString("password"),
	}

	result := repo.Login(loginForm)

	if result == "gagal login" {
		fmt.Println("[LOGIN ERROR] Email atau Password salah")
		c.Data["messageError"] = "Email atau Password salah"
		c.Redirect("/", 302)
		return
	}

	// kalau berhasil
	c.SetSession("controllerSession", loginForm.Email) // bisa simpan ID/Email user
	c.Data["messageSuccess"] = "Berhasil login"
	fmt.Println("[LOGIN SUCCESS] User:", loginForm.Email)

	c.Redirect("/dashboard", 302)
}
