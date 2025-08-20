package api

import (
	"github.com/astaxie/beego"
	"wms/repo/goods"
	Units "wms/repo/units"
	"wms/models"
	"fmt"
	"strconv"
	"math/rand"
	"time"
	
)


type MainController struct {
	beego.Controller
}


func (c *MainController) All() {
	//Check Session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	goodsRepo := repo.GoodsRepo{}

	flash := beego.ReadFromRequest(&c.Controller)
	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		c.Data["flash"] = ok
	}

	result, err := goodsRepo.GetAll()
	unitsRepo := Units.UnitRepo{}
	listUnits, errGetListUnits := unitsRepo.GetAll()

	if errGetListUnits != nil {
		c.Data["errormsg"] = errGetListUnits
		c.Layout = "template.html"
		c.TplName = "incoming.html"
	}

	if err != nil {
		c.Data["pesan"] = err
		c.Layout = "template.html"
		c.TplName = "incoming.html"
	}
	
	c.Data["incoming"] = result
	c.Data["list"] = listUnits
	c.Layout = "template.html"
	c.TplName = "incoming.html"
}


func (c *MainController) Add() {
	t := time.Now()
	goodsRepo := repo.GoodsRepo{}
	goodsModel :=  models.Incoming{}
	newTrx := generateTrxID()
	//mau nampung data dari form
	goodsModel.TransaksiId 	= newTrx
	goodsModel.Tanggal 		= t.Format("2006/01/02 15:04:05")
	goodsModel.Lokasi 		= c.GetString("lokasi")
	goodsModel.KodeBarang 	= c.GetString("kode")
	goodsModel.NamaBarang 	= c.GetString("nama")
	goodsModel.Satuan 		= c.GetString("satuan")
	
	i, _ := strconv.Atoi(c.GetString("jumlah"))
	goodsModel.Jumlah = i

	result, err := goodsRepo.Add(goodsModel)

	if err != nil {
		fmt.Println(result)
		c.Data["messageSuccess"]= "Gagal menambahkan barang"
		c.Layout 				= "template.html"
		c.TplName 				= "incoming.html"
	}
	
	fmt.Println(result)
	c.Data["messageSuccess"] = "Berhasil menambahkan barang"
	c.Redirect("/incoming-goods",302)
}


//Incoming Detail
func (c *MainController) Detail() {
	//Check Session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	goodsRepo := repo.GoodsRepo{}

	id := c.Ctx.Input.Param(":id")
	i, _ := strconv.Atoi(id)
	result, err := goodsRepo.GetById(i)

	if err != nil {
		fmt.Println("Error Repo => ",err)
		c.Data["errormsg"] = err
		c.Layout = "template.html"
		c.TplName = "form-detail.html"
	}

	//Panggil repo get all satuan
	unitsRepo := Units.UnitRepo{}
	listUnits, errGetListUnits := unitsRepo.GetAll()

	if err != nil {
		fmt.Println("Error Repo => ",errGetListUnits)
		c.Data["errormsg"] = errGetListUnits
		c.Layout = "template.html"
		c.TplName = "form-detail.html"
	}
	
	c.Data["list"] = listUnits
	c.Data["detail"] = result
	c.Layout = "template.html"
	c.TplName = "form-detail.html"
}

func (c *MainController) DetailProccess() {
	//Check Session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	goodsRepo := repo.GoodsRepo{}

	id := c.Ctx.Input.Param(":id")
	i, _ := strconv.Atoi(id)
	result, err := goodsRepo.GetById(i)

	if err != nil {
		fmt.Println("Error Repo => ",err)
		c.Data["errormsg"] = err
		c.Layout = "template.html"
		c.TplName = "form-detail.html"
	}

	//Panggil repo get all satuan
	unitsRepo := Units.UnitRepo{}
	listUnits, errGetListUnits := unitsRepo.GetAll()

	if err != nil {
		fmt.Println("Error Repo => ",errGetListUnits)
		c.Data["errormsg"] = errGetListUnits
		c.Layout = "template.html"
		c.TplName = "form-proses.html"
	}
	
	c.Data["list"] = listUnits
	c.Data["detail"] = result
	c.Layout = "template.html"
	c.TplName = "form-proses.html"
}

//Update Detail
func (c *MainController) Update() {
	goodsModel := models.Incoming{}
	goodsRepo := repo.GoodsRepo{}

	id := c.Ctx.Input.Param(":id")
	i, _ := strconv.Atoi(id)
	goodsModel.Id = i
	goodsModel.Lokasi 		= c.GetString("lokasi")
	goodsModel.KodeBarang 	= c.GetString("kode")
	goodsModel.NamaBarang 	= c.GetString("nama")
	goodsModel.Satuan 		= c.GetString("namaSatuan")
	
	qty, _ := strconv.Atoi(c.GetString("jumlah"))
	goodsModel.Jumlah = qty
	fmt.Println(goodsModel)
	result, err := goodsRepo.Update(goodsModel)
	if err != nil {
		c.Data["message"] = err
		c.Redirect("/incoming-goods",302)
	}
	c.Data["message"] = result
	c.Redirect("/incoming-goods",302)
}


func (c *MainController) Delete() {
	goodsRepo := repo.GoodsRepo{}

	id := c.Ctx.Input.Param(":id")
	i, _ := strconv.Atoi(id)

	num, err := goodsRepo.Delete(i)

	if err != nil {
		c.Data["errormsg"] = err
		c.Redirect("/incoming-goods",302)
	}
	c.Data["pesan"] = string(num)
	c.Redirect("/incoming-goods",302)
}


//Helper Function
func generateTrxID() string{
	rand.Seed(time.Now().Unix())
    
	pattern := "TRX"
	num1 := randomWithRange(100, 500)
	num2 := randomWithRange(501, 999)

	s1 := strconv.Itoa(num1)
	s2 := strconv.Itoa(num2)

	generate := pattern + "-" + s1 + "-" + s2
	return generate
}


func randomWithRange(min, max int) int {
    var value = rand.Int() % (max - min + 1) + min
    return value
}





