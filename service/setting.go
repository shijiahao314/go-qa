package service

import (
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"go.uber.org/zap"
)

type SettingService struct{}

func (ss *SettingService) UpdateSetting(userid uint64, setting model.UserSetting) error {
	tx := global.DB.Begin()
	var cnt int64
	if err := global.DB.Model(&model.UserSetting{}).Where("userid = ?", userid).Count(&cnt).Error; err != nil {
		tx.Rollback()
		global.Logger.Error("UpdateSetting failed", zap.Error(err))
		return err
	}
	if err := global.DB.Model(&model.UserSetting{}).Where("userid = ?", userid).Updates(setting).Error; err != nil {
		tx.Rollback()
		global.Logger.Error("UpdateSetting failed", zap.Error(err))
		return err
	}
	tx.Commit()
	return nil
}
