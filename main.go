package main

import (
	"fmt"
	"os/exec"

	_ "wms/routers"

	"github.com/astaxie/beego"
)

func openChrome(url string) error {
	chromePath := `C:\Program Files\Google\Chrome\Application\chrome.exe`
	cmd := exec.Command(chromePath, url)
	return cmd.Start()
}

func main() {
	// Set port Beego sesuai URL
	beego.BConfig.Listen.HTTPPort = 8081

	// Konfigurasi dev mode (opsional)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.Session.SessionOn = true
	}

	// Buka Chrome otomatis
	url := "http://localhost:8081"
	if err := openChrome(url); err != nil {
		fmt.Println("Gagal membuka Chrome:", err)
	} else {
		fmt.Println("Chrome berhasil dibuka!")
	}

	// Jalankan server Beego
	beego.Run()
}
