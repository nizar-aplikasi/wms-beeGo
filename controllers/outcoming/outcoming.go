package api

import (
	"github.com/astaxie/beego"
	Goods "wms/repo/goods"
	"wms/models"
	"strconv"
	"time"
	"fmt"
	// "github.com/SebastiaanKlippert/go-wkhtmltopdf"
    // "log"
    // "os"
)


type OutController struct {
	beego.Controller
}

func (c *OutController) All() {
	//Check Session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	repoOutcoming := Goods.GoodsRepo{}
	result, err := repoOutcoming.AllOut()
	if err != nil {
		c.Data["errormsg"] = err
		c.Layout = "template.html"
		c.TplName = "outcoming.html" //buat load halaman
	}
	fmt.Println(result)
	c.Data["list"] = result
		c.Layout = "template.html"
		c.TplName = "outcoming.html" //buat load halaman
}

func (c *OutController) Detail() {
	//Check Session
	v := c.GetSession("controllerSession")
    if v == nil {
        c.Redirect("/", 302)
    }
	repoOutcoming := Goods.GoodsRepo{}
	i, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	
	fmt.Println(i)
	result, err := repoOutcoming.OutById(i)
	if err != nil {
		c.Data["errormsg"] = err
		c.Layout = "template.html"
		c.TplName = "form-out-detail.html"
	}
	fmt.Println(result,"masuk sini")
	c.Data["detail"] = result
	c.Layout = "template.html"
	c.TplName = "form-out-detail.html"
}

func (c *OutController) Update() {
	repoOutcoming := Goods.GoodsRepo{}
	outModel := models.Outcoming{}
	i, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	outModel.Id = i
	outModel.TransaksiId = c.GetString("trxId")
	outModel.TanggalMasuk = c.GetString("tanggalmasuk")
	outModel.TanggalKeluar = c.GetString("tanggalkeluar")
	outModel.Lokasi = c.GetString("lokasi")
	outModel.KodeBarang = c.GetString("kode")
	outModel.NamaBarang = c.GetString("nama")
	iJumlah, _ := strconv.Atoi(c.GetString("jumlah"))
	outModel.Jumlah = iJumlah
	
	fmt.Println(i, "=> query param id")
	result, err := repoOutcoming.OutUpdate(outModel)
	if err != nil {
		c.Data["errormsg"] = err
		c.Redirect("/outcoming-goods", 302)
	}
	fmt.Println(result,"masuk sini")
	c.Data["detail"] = result
	c.Layout = "template.html"
	c.Redirect("/outcoming-goods", 302)
}

func (c *OutController) OutProccess() {
	t := time.Now()
	goodsRepo := Goods.GoodsRepo{}
	outModel :=  models.Outcoming{}
	//mau nampung data dari form
	getID, _ := strconv.Atoi(c.GetString("id"))
	outModel.Id 	= getID
	outModel.TransaksiId 	= c.GetString("trxId")
	outModel.TanggalMasuk 	= c.GetString("tanggal")
	outModel.TanggalKeluar 	= t.Format("2006/01/02 15:04:05")
	outModel.Lokasi 		= c.GetString("lokasi")
	outModel.KodeBarang 	= c.GetString("kode")
	outModel.NamaBarang 	= c.GetString("nama")
	outModel.Satuan 		= c.GetString("satuan")
	
	i, _ := strconv.Atoi(c.GetString("jumlah"))
	outModel.Jumlah = i
	
	result, err := goodsRepo.OutProccess(outModel)

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


