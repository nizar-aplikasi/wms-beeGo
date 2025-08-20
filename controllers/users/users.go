package api

import (
	"github.com/astaxie/beego"
	Users "wms/repo/users"
	"wms/models"
	"fmt"
	"strconv"
)


type UserController struct {
	beego.Controller
}


func (c *UserController) All() {
	//Check session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	usersRepo := Users.UsersRepo{}
	result, errGetUsers := usersRepo.All()

	if errGetUsers != nil {
		c.Data["errormsg"] = errGetUsers
		c.Layout = "template.html"
		c.TplName = "users.html"
	}

	if errGetUsers != nil {
		c.Data["pesan"] = errGetUsers	
		c.Layout = "template.html"
		c.TplName = "users.html"
	}
	
	fmt.Println(result)
	c.Data["list"] = result
	c.Layout = "template.html"
	c.TplName = "users.html"
}


func (c *UserController) Add() {
	userRepo := Users.UsersRepo{}
	userModel :=  models.Users{}
	//mau nampung data dari form
	userModel.Username 		= c.GetString("username")
	userModel.Email 	= c.GetString("email")
	userModel.Password 	= c.GetString("password")
	i, _ := strconv.Atoi(c.GetString("role"))
	userModel.Role = i
	fmt.Println(i,"debug user role")
	result, err := userRepo.Add(userModel)

	if err != nil {
		fmt.Println(result)
		c.Data["messageSuccess"]= "Gagal menambahkan user"
		c.Layout 				= "template.html"
		c.TplName 				= "users.html"
	}
	
	fmt.Println(result)
	c.Data["messageSuccess"] = "Berhasil menambahkan user"
	c.Redirect("/users",302)
}


// //Incoming Detail
// func (c *MainController) Detail() {
// 	goodsRepo := repo.GoodsRepo{}

// 	id := c.Ctx.Input.Param(":id")
// 	i, _ := strconv.Atoi(id)
// 	result, err := goodsRepo.GetById(i)

// 	if err != nil {
// 		fmt.Println("Error Repo => ",err)
// 		c.Data["errormsg"] = err
// 		c.Layout = "template.html"
// 		c.TplName = "form-detail.html"
// 	}

// 	//Panggil repo get all satuan
// 	unitsRepo := Units.UnitRepo{}
// 	listUnits, errGetListUnits := unitsRepo.GetAll()

// 	if err != nil {
// 		fmt.Println("Error Repo => ",errGetListUnits)
// 		c.Data["errormsg"] = errGetListUnits
// 		c.Layout = "template.html"
// 		c.TplName = "form-detail.html"
// 	}
	
// 	c.Data["list"] = listUnits
// 	c.Data["detail"] = result
// 	c.Layout = "template.html"
// 	c.TplName = "form-detail.html"
// }

// func (c *MainController) DetailProccess() {
// 	goodsRepo := repo.GoodsRepo{}

// 	id := c.Ctx.Input.Param(":id")
// 	i, _ := strconv.Atoi(id)
// 	result, err := goodsRepo.GetById(i)

// 	if err != nil {
// 		fmt.Println("Error Repo => ",err)
// 		c.Data["errormsg"] = err
// 		c.Layout = "template.html"
// 		c.TplName = "form-detail.html"
// 	}

// 	//Panggil repo get all satuan
// 	unitsRepo := Units.UnitRepo{}
// 	listUnits, errGetListUnits := unitsRepo.GetAll()

// 	if err != nil {
// 		fmt.Println("Error Repo => ",errGetListUnits)
// 		c.Data["errormsg"] = errGetListUnits
// 		c.Layout = "template.html"
// 		c.TplName = "form-proses.html"
// 	}
	
// 	c.Data["list"] = listUnits
// 	c.Data["detail"] = result
// 	c.Layout = "template.html"
// 	c.TplName = "form-proses.html"
// }

// //Update Detail
// func (c *MainController) Update() {
// 	goodsModel := models.Incoming{}
// 	goodsRepo := repo.GoodsRepo{}

// 	id := c.Ctx.Input.Param(":id")
// 	i, _ := strconv.Atoi(id)
// 	goodsModel.Id = i
// 	goodsModel.Lokasi 		= c.GetString("lokasi")
// 	goodsModel.KodeBarang 	= c.GetString("kode")
// 	goodsModel.NamaBarang 	= c.GetString("nama")
// 	goodsModel.Satuan 		= c.GetString("namaSatuan")
	
// 	qty, _ := strconv.Atoi(c.GetString("jumlah"))
// 	goodsModel.Jumlah = qty
// 	fmt.Println(goodsModel)
// 	result, err := goodsRepo.Update(goodsModel)
// 	if err != nil {
// 		c.Data["message"] = err
// 		c.Redirect("/incoming-goods",302)
// 	}
// 	c.Data["message"] = result
// 	c.Redirect("/incoming-goods",302)
// }


// func (c *MainController) Delete() {
// 	goodsRepo := repo.GoodsRepo{}

// 	id := c.Ctx.Input.Param(":id")
// 	i, _ := strconv.Atoi(id)

// 	num, err := goodsRepo.Delete(i)

// 	if err != nil {
// 		c.Data["errormsg"] = err
// 		c.Redirect("/incoming-goods",302)
// 	}
// 	c.Data["pesan"] = string(num)
// 	c.Redirect("/incoming-goods",302)
// }


// //Helper Function
// func generateTrxID() string{
// 	rand.Seed(time.Now().Unix())
    
// 	pattern := "TRX"
// 	num1 := randomWithRange(100, 500)
// 	num2 := randomWithRange(501, 999)

// 	s1 := strconv.Itoa(num1)
// 	s2 := strconv.Itoa(num2)

// 	generate := pattern + "-" + s1 + "-" + s2
// 	return generate
// }


// func randomWithRange(min, max int) int {
//     var value = rand.Int() % (max - min + 1) + min
//     return value
// }





