package service

import (
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"go.uber.org/zap"
)

type SettingService struct{}

func (ss *SettingService) UpdateSetting(userid uint64, settingDTO model.UserSettingDTO) error {
	tx := global.DB.Begin()
	var cnt int64
	if err := global.DB.Model(&model.UserSetting{}).Where("user_id = ?", userid).Count(&cnt).Error; err != nil {
		tx.Rollback()
		global.Logger.Error("UpdateSetting failed", zap.Error(err))
		return err
	}
	setting := &model.UserSetting{
		UserID:       userid,
		OpenaiApiKey: settingDTO.OpenaiApiKey,
		ChatModel:    settingDTO.ChatModel,
		TestMode:     settingDTO.TestMode,
	}
	if cnt == 0 {
		if err := global.DB.Create(&setting).Error; err != nil {
			tx.Rollback()
			global.Logger.Error("UpdateSetting failed", zap.Error(err))
			return err
		}
		tx.Commit()
		return nil
	}
	if err := global.DB.Model(&model.UserSetting{}).Where("user_id = ?", userid).Updates(settingDTO).Error; err != nil {
		tx.Rollback()
		global.Logger.Error("UpdateSetting failed", zap.Error(err))
		return err
	}
	tx.Commit()
	return nil
}

func (ss *SettingService) GetSetting(userid uint64) (model.UserSetting, error) {
	setting := model.UserSetting{}
	tx := global.DB.Begin()
	if err := tx.Model(&model.UserSetting{}).Where("user_id = ?", userid).First(&setting).Error; err != nil {
		tx.Rollback()
		global.Logger.Error("GetSetting failed", zap.Error(err))
		return setting, err
	}
	tx.Commit()
	return setting, nil
}
