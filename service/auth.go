package service

import (
	"fmt"

	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func (as *AuthService) Login(username, password string) (model.User, error) {
	var user model.User
	if err := global.DB.Model(&model.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return user, fmt.Errorf("failed to find user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		user.Password = ""
		return user, fmt.Errorf("incorrect password: %w", err)
	}

	user.Password = ""
	return user, nil
}

func (as *AuthService) UserIDInDatabase(userID uint64) (bool, error) {
	var user model.User
	if err := global.DB.Model(&model.User{}).Where("user_id = ?", userID).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (as *AuthService) UserHasPermission(userid uint64) (bool, error) {
	var user model.User
	if err := global.DB.Model(&model.User{}).Where("user_id = ?", userid).First(&user).Error; err != nil {
		return false, err
	}
	if user.Role == global.RoleAdmin {
		return true, nil
	}
	return false, nil
}
