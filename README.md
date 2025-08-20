# wms-beeGo

Warehouse Management System (WMS) menggunakan **Go Lang** dengan **Beego Framework** dan frontend **AdminLTE**.  
Project ini dibuat untuk belajar Clean Architecture di Go, integrasi database PostgreSQL, dan implementasi REST API.

---

## 🚀 Fitur

### 1. Autentikasi
- Login: `/login` (POST)  
- Logout: `/logout` (GET)  
- Auth API: `/auth` (POST)

### 2. Dashboard
- Halaman utama dashboard: `/dashboard` (GET)

### 3. Manajemen Barang
#### Incoming Goods (Barang Masuk)
- Daftar barang masuk: `/incoming-goods` (GET)  
- Tambah barang masuk: `/incoming-goods/add` (POST)  
- Detail barang masuk: `/incoming-goods/detail/:id` (GET)  
- Update barang masuk: `/incoming-goods/update/:id` (POST)  
- Proses barang keluar: `/incoming-goods/proccess/:id` (POST)  
- Hapus barang masuk: `/incoming-goods/delete/:id`  

#### Outcoming Goods (Barang Keluar)
- Daftar barang keluar: `/outcoming-goods` (GET)  
- Detail barang keluar: `/outcoming-goods/detail/:id` (GET)  
- Update barang keluar: `/outcoming-goods/update/:id` (POST)  

### 4. Manajemen Unit
- Daftar unit: `/units` (GET)  
- Tambah unit: `/units/add` (POST)  
- Detail unit: `/units/detail/:id` (GET)  
- Update unit: `/units/update/:id` (POST)  
- Hapus unit: `/units/delete/:id`  

### 5. Manajemen User
- Daftar user: `/users` (GET)  
- Tambah user: `/users/add` (POST)  
- Hapus user: `/users/delete/:id`  

### 6. Reporting
- Lihat laporan: `/reporting` (GET)  
- Export laporan: `/reporting/export` (GET)

## 🛠 Teknologi

- **Backend**: Go Lang, Beego Framework  
- **Database**: MySQL 
- **Frontend**: AdminLTE, jQuery, Bootstrap 4  
- **Tools**: Git, GitHub  

---
