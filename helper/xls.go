package helper

import(
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	RepoReporting "wms/repo/reporting"
	"wms/models"
)

type HelperModule struct{}

var Xlsx *excelize.File

func (hm *HelperModule) ExportToExcell() bool{
	repo := RepoReporting.ReportRepo{}
	result, err := repo.All()
	if err != nil {
		fmt.Println("Error saat mengambil data reporting...")
		return false
	}
	Xlsx = excelize.NewFile()

	sheet1Name := "Sheet1"
	Xlsx.SetSheetName(Xlsx.GetSheetName(1), sheet1Name)

	//prepare column
	Xlsx.SetCellValue(sheet1Name, "A1", "Nama Barang")
	Xlsx.SetCellValue(sheet1Name, "B1", "Lokasi Gudang")
	Xlsx.SetCellValue(sheet1Name, "C1", "Kode Satuan")
	Xlsx.SetCellValue(sheet1Name, "D1", "Jumlah Barang Masuk")
	Xlsx.SetCellValue(sheet1Name, "E1", "Jumlah Barang Keluar")
	Xlsx.SetCellValue(sheet1Name, "F1", "Sisa Barang")

	errFilterXLS := Xlsx.AutoFilter(sheet1Name, "A1", "F1", "")
	if errFilterXLS != nil {
		fmt.Println(errFilterXLS)
		return false
	}


	y := make([]interface{}, len(result))
	for i, each := range result {
    	y[i] = each
	}
	

	for i, eachInterface := range y {
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), eachInterface.(models.Reporting).NamaBarang)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), eachInterface.(models.Reporting).Lokasi)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), eachInterface.(models.Reporting).Kode)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("D%d", i+2), eachInterface.(models.Reporting).JumlahMasuk)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("E%d", i+2), eachInterface.(models.Reporting).JumlahKeluar)
    	Xlsx.SetCellValue(sheet1Name, fmt.Sprintf("F%d", i+2), eachInterface.(models.Reporting).Sisa)
	}

	errCreateXLS := Xlsx.SaveAs("./static/laporan-keluar-masuk-barang.xlsx")
	if errCreateXLS != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Created xlsx success")
	return true
}


