package models

type User struct {
    Id          int     `json:"id"`
    Username    string  `json:"username"`
    Email       string  `json:"email"`
    Password    string  `json:"password"`
    Role        int     `json:"role"`
}

type Unit struct {
  Id          int     `json:"id"`
  Code        string  `json:"code"`
  Name        string  `json:"name"`
}

type Incoming struct {
  Id            int     `json:"id"`
  TransaksiId   string  `json:"transaksiid"`
  Tanggal       string  `json:"tanggal"`
  Lokasi        string  `json:"lokasi"`
  KodeBarang    string  `json:"kodebarang"`
  NamaBarang    string  `json:"namabarang"`
  Satuan        string  `json:"satuan"`
  Jumlah        int     `json:"jumlah"`
}


type Outcoming struct {
  Id            int     `json:"id"`
  TransaksiId   string  `json:"transaksi_id"`
  TanggalMasuk  string  `json:"tanggalmasuk"`
  TanggalKeluar string  `json:"tanggalkeluar"`
  Lokasi        string  `json:"lokasi"`
  KodeBarang    string  `json:"kodebarang"`
  NamaBarang    string  `json:"namabarang"`
  Satuan        string  `json:"satuan"`
  Jumlah        int     `json:"jumlah"`
}

type Users struct {
  Id            int        `json:"id"`
  Username      string     `json:"username"`
  Email         string     `json:"email"`
  Password      string     `json:"password"`
  Role          int        `json:"role"`
  LastLogin     string     `json:"lastLogin"`
}

type Units struct {
  Id     int       `json:"id"`
  Code   string    `json:"kode"`
  Name   string    `json:"nama"`
}

type MaxIncomingGoods struct{
		Max int
}

type Reporting struct{
  NamaBarang    string  `json:"namabarang"`
  Lokasi    string  `json:"lokasi"`
  Kode    string  `json:"kode"`
  Satuan    string  `json:"satuan"`
  JumlahMasuk   int     `json:"jumlahmasuk"`
  JumlahKeluar int     `json:"jumlahkeluar"`
  Sisa      int     `json:"sisa"`
}

type LoginForm struct{
	Email string
  Password string
}

type LoginResponse struct{
  Username string
	Email string
  Password string
  Role string
}

