package repo

import "wms/database"
import "github.com/astaxie/beego"

type RepoDashboard struct{
	beego.Controller
}


func (repo *RepoDashboard) GetTotalIncomingGoods() (int, error){

    var totalBarangMasuk int
    if result := database.DB.Table("incoming").Count(&totalBarangMasuk); result.Error != nil {
		return 0, result.Error
	}
	return int(totalBarangMasuk), nil
}	

func (repo *RepoDashboard) GetTotalOutcomingGoods() (result int, err error){
	var totalBarangKeluar int
    if result := database.DB.Table("outcoming").Count(&totalBarangKeluar); result.Error != nil {
		return 0, result.Error
	}
	return totalBarangKeluar, nil
}	

func (repo *RepoDashboard) GetTotalUser() (result int, err error){
	var totalUser int
    if result := database.DB.Table("users").Count(&totalUser); result.Error != nil {
		return 0, result.Error
	}
	return totalUser, nil
}	

func (repo *RepoDashboard) Max() (float64, string){
	var result float64
		row := database.DB.Table("incoming").Select("max(jumlah)").Row()
		row.Scan(&result)
		return result, "Data max jumlah barang berhasil ditampikan"
}



