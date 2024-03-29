package service

import (
	"fmt"

	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (us *UserService) AddUser(u model.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt generate error: %w", err)
	}

	user := model.User{
		UserID:      utils.GetSnowflakeID(),
		AccountType: u.AccountType,
		Username:    u.Username,
		Password:    string(hashedPass),
		Role:        u.Role,
	}

	var cnt int64
	tx := global.DB.Begin()
	if err := tx.Model(&model.User{}).Where("username = ?", u.Username).Count(&cnt).Error; err != nil {
		tx.Rollback()
		return err
	}
	if cnt > 0 {
		return fmt.Errorf("user %s already exists", u.Username)
	}
	if err := tx.Model(&model.User{}).Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	// casbin: add policy
	global.Enforcer.AddGroupingPolicy(user.Username, string(user.Role))
	tx.Commit()
	return nil
}

func (us *UserService) DeleteUser(id uint64) error {
	if err := global.DB.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(u model.User) error {
	var user model.User
	tx := global.DB.Begin()

	if err := tx.Model(&model.User{}).Where("user_id = ?", u.UserID).Take(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt generate error: %w", err)
	}

	// update
	user.Password = string(hashedPass)
	if len(u.Role) > 0 {
		user.Role = u.Role
	}

	if err := tx.Model(&model.User{}).Where("user_id = ?", u.UserID).Save(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (us *UserService) UsernameExists(username string) (bool, error) {
	var cnt int64
	if err := global.DB.Model(&model.User{}).Where("username = ?", username).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (us *UserService) GithubUserExists(id uint64) (bool, error) {
	var cnt int64
	if err := global.DB.Model(&model.GithubUser{}).Where("user_id = ?", id).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (us *UserService) GetUser(id uint64) (*model.User, error) {
	var user model.User
	if err := global.DB.Model(&model.User{}).Where("user_id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) GetUsers(page, size int) ([]model.User, int64, error) {
	users := make([]model.User, 0)
	var cnt int64

	tx := global.DB.Begin()

	if err := tx.Model(&model.User{}).Offset((page - 1) * size).Limit(size).Find(&users).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	if err := tx.Model(&model.User{}).Count(&cnt).Error; err != nil {
		tx.Rollback()
		return nil, 0, err
	}

	for _, u := range users {
		u.Password = ""
	}

	tx.Commit()
	return users, cnt, nil
}
