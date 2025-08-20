package repo

import "wms/models"
import "wms/database"
import "github.com/astaxie/beego"
import "fmt"

type GoodsRepo struct{
	beego.Controller
}

//Buat kondisi where jumlah tidak != 0
func (repo *GoodsRepo) GetAll() (result []models.Incoming ,err error){


    var barangMasuk []models.Incoming
	database.DB.Raw("select * from incoming where jumlah > 0 order by id desc").Scan(&barangMasuk)
	// if result := database.DB.Find(&barangMasuk); result.Error != nil {
	// 	return barangMasuk, result.Error
	// }

	return barangMasuk, err
}	


func (repo *GoodsRepo) AllOut() ([]models.Outcoming , error){

	var barangKeluar []models.Outcoming
	database.DB.Raw("select * from outcoming order by id desc").Scan(&barangKeluar)
	// if result := database.DB.Find(&barangKeluar); result.Error != nil {
	// 	return barangKeluar, result.Error
	// }
	return barangKeluar, nil
}	

func (repo *GoodsRepo) GetById(id int) (result models.Incoming ,err error){
	
	DetailBarang := models.Incoming{}

	if result := database.DB.Where("id = ?", id).Find(&DetailBarang); result.Error != nil {
		return DetailBarang , result.Error
	}

	return DetailBarang , nil
}	


func (repo *GoodsRepo) OutById(id int) (models.Outcoming , error){
	
	OutDetail := models.Outcoming{}
	if result := database.DB.Where("id = ?", id).Find(&OutDetail); result.Error != nil {
		return OutDetail , result.Error
	}

	return OutDetail , nil
}

func (repo *GoodsRepo) Delete(id int) (string ,error){
	DeleteBarang := models.Incoming{}
	DeleteBarang.Id = id
	if result := database.DB.Delete(&DeleteBarang); result.Error != nil {
		return "Hapus barang gagal", result.Error
	}
	
	return "Delete barang berhasil", nil
}	

func (repo *GoodsRepo) Add(form models.Incoming) ( bool, error){
	
	database.DB.Create(&form)
	return true, nil
}	

func (repo *GoodsRepo) OutProccess(form models.Outcoming) ( string, error){
	if result := database.DB.Create(&form); result.Error != nil {
		return "proses barang gagal", result.Error
	}
	return "proses barang berhasil", nil
}	

func (repo *GoodsRepo) Update(form models.Incoming) (string, error){
	UpdateBarang := models.Incoming{}
	if result := database.DB.Model(&UpdateBarang).Updates(form); result.Error != nil {
		return "Update gagal", result.Error
	}
	return "Update data berhasil", nil
}

func (repo *GoodsRepo) OutUpdate(form models.Outcoming) (string, error){
	fmt.Println(form)
	outUpdate := models.Outcoming{}
	if result := database.DB.Model(&outUpdate).Updates(form); result.Error != nil {
		return "Update gagal", result.Error
	}
	return "Update data berhasil", nil
}












