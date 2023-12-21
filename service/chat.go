package service

import (
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
)

type ChatService struct{}

// Add ChatInfo
func (cs *ChatService) AddChatInfo(chatInfo model.ChatInfo) error {
	tx := global.DB.Begin()

	if err := tx.Create(&chatInfo).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// 查找该用户下所有的ChatInfo
func (cs *ChatService) GetChatInfos(userId uint64) ([]model.ChatInfo, error) {
	chatInfos := make([]model.ChatInfo, 0)

	tx := global.DB.Begin()

	// Find查找UserID所有ChatInfo
	if err := tx.Model(&model.ChatInfo{}).Where(&model.ChatInfo{UserID: userId}).Find(&chatInfos).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return chatInfos, nil
}
