package repo

import (
	"wms/database"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"wms/models"
)

type AuthRepo struct{
	beego.Controller
}



func(repo *AuthRepo) Login(form models.LoginForm) (result models.LoginResponse, text string){
	//Logicnya cari di table users yg emailnya == input email, kalo ada 
	//maka ambil passwordnya lalu di compare
	userModel := models.LoginResponse{}
	err := database.DB.Raw("SELECT username, email, password, role FROM users WHERE email = ?", form.Email).Scan(&userModel)
	if err != nil {
		return userModel, "Login Gagal"
	}
	match := repo.CheckPassword(userModel.Password, form.Password)
		if !match{
			beego.Info("Login Gagal")
		} else {
			return userModel, "Login Gagal"
		}
	return userModel, "Login Berhasil"
}

func (repo *AuthRepo) CheckPassword(hashPwd string, plainPwd string) bool{
	/*
		PENJELASAN
		Fungsi ini membandingkan antara password hash yang ada di db dengan inputan user, returnnya adalah false or true.
		Fungsi ini menerima 2 argumen yaitu string password yg sudah di hash, dan string password dari form input
	*/

	//Kedua argumen dirubah jadi byte code
	byteHash := []byte(hashPwd)
	bytePlain := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}