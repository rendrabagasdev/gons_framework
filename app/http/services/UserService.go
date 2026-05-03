package services

import (
	"go-framework/app/models"
	"log/slog"

	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB `container:"type"`
}

/*
Constructor user service
@return *UserService
*/
func NewUserService() *UserService {
	scv := &UserService{}
	if err := container.Fill(scv); err != nil {
		slog.Error("Gons: UserService injection error", "error", err)
	}
	return scv
}

// auto registrasi ke service registry
func init() {
	RegisterService(func() error {
		return container.Singleton(func() *UserService {
			return NewUserService()
		})
	})
}

// Get All Users
func (service *UserService) GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	err := service.DB.Find(&users).Error
	return users, err
}
