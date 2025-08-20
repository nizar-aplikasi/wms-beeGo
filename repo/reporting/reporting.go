package repo

import "wms/models"
import "wms/database"
import "github.com/astaxie/beego"


type ReportRepo struct{
	beego.Controller
}

//Buat kondisi where jumlah tidak != 0
func (repo *ReportRepo) All() (result []models.Reporting ,err error){


    var laporanBarang []models.Reporting
	database.DB.Raw("select * from v_in_out_reporting").Scan(&laporanBarang)
	// if result := database.DB.Find(&barangMasuk); result.Error != nil {
	// 	return barangMasuk, result.Error
	// }
	return laporanBarang, err
}	















