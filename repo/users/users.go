package repo

import (
	"fmt"
	"wms/database"
	"wms/models"

	"golang.org/x/crypto/bcrypt"
)

// UsersRepo berisi method untuk akses data Users
type UsersRepo struct{}

// Ambil semua user
func (repo *UsersRepo) All() ([]models.Users, error) {
	var users []models.Users
	if result := database.DB.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Tambah user baru dengan password ter-hash
func (repo *UsersRepo) Add(form models.Users) (string, error) {
	hashedPassword, err := repo.HashingWithSalt(form.Password)
	if err != nil {
		return "", err
	}
	form.Password = hashedPassword

	if result := database.DB.Create(&form); result.Error != nil {
		return "", result.Error
	}

	return "berhasil menambahkan users", nil
}

// Login cek ke DB
func (repo *UsersRepo) Login(form models.LoginForm) string {
	var user models.Users

	// cek user by email
	if err := database.DB.Where("email = ?", form.Email).First(&user).Error; err != nil {
		fmt.Println("User tidak ditemukan:", err)
		return "gagal login"
	}

	// validasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		fmt.Println("Password salah:", err)
		return "gagal login"
	}

	return "berhasil login"
}

// Hashing password dengan bcrypt
func (repo *UsersRepo) HashingWithSalt(password string) (string, error) {
	hashAndSalt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("proses hashing error: %v", err)
	}
	return string(hashAndSalt), nil
}
