package repo

import(
	"wms/database"
	"wms/models"
	"github.com/astaxie/beego"
)

type UnitRepo struct{
	beego.Controller
}

func (repo *UnitRepo) GetAll() (result []models.Units ,err error){
	
    var units []models.Units
	
	if result := database.DB.Find(&units);  result.Error != nil {
		return units, result.Error
	}
	
	return units, nil
}

func (repo *UnitRepo) Add(form models.Units) (string, error){
	
	if result := database.DB.Create(&form);  result.Error != nil {
		return "gagal menambahkan satuan barang ", result.Error
	}
	
	return "satuan barang berhasil ditambahkan", nil
}


func (repo *UnitRepo) GetById(id int) (models.Units ,error){
	DetailUnit := models.Units{}

	if result := database.DB.Where("id = ?", id).Find(&DetailUnit); result.Error != nil {
		return DetailUnit , result.Error
	}

	return DetailUnit , nil
}	

func (repo *UnitRepo) Update(form models.Units) (string, error){
	updateUnit := models.Units{}
	if result := database.DB.Model(&updateUnit).Updates(form); result.Error != nil {
		return "Update unit gagal", result.Error
	}
	return "Update unit berhasil", nil
}

func (repo *UnitRepo) Delete(id int) (string ,error){
	deleteUnit := models.Units{}
	deleteUnit.Id = id
	if result := database.DB.Delete(&deleteUnit); result.Error != nil {
		return "Hapus satuan gagal", result.Error
	}
	
	return "Delete satuan berhasil", nil
}	



