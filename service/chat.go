package service

import (
	"context"
	"fmt"

	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
)

type ChatService struct{}

// Check userId是否与ChatInfo.UserID匹配
func (cs *ChatService) CheckUser(userId uint64, id uint) error {
	var chatInfo model.ChatInfo
	if err := global.DB.Model(&model.ChatInfo{}).Where(&model.ChatInfo{ID: id}).Take(&chatInfo).Error; err != nil {
		return err
	}
	if userId != chatInfo.UserID {
		return fmt.Errorf("user id [%d] does not match chat info user id [%d]", userId, chatInfo.UserID)
	}
	return nil
	context.Context
}

// Add ChatInfo
func (cs *ChatService) AddChatInfo(chatInfo model.ChatInfo) error {
	if err := global.DB.Create(&chatInfo).Error; err != nil {
		return err
	}

	return nil
}

// Delete ChatInfo by id
func (cs *ChatService) DeleteChatInfo(id uint) error {
	tx := global.DB.Begin()

	// 删除ChatInfo下的ChatCard
	if err := global.DB.Where(&model.ChatCard{ChatInfoID: id}).Delete(&model.ChatCard{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除ChatInfo
	if err := global.DB.Delete(&model.ChatInfo{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// Update ChatInfo
func (cs *ChatService) UpdateChatInfo(chatInfo model.ChatInfo) error {
	if err := global.DB.Model(&model.ChatInfo{}).Where(&model.ChatInfo{ID: chatInfo.ID}).Save(&chatInfo).Error; err != nil {
		return err
	}

	return nil
}

// Get ChatInfo by id
func (cs *ChatService) GetChatInfo(id uint) (model.ChatInfo, error) {
	chatInfo := model.ChatInfo{}
	if err := global.DB.Model(&model.ChatInfo{}).Where(&model.ChatInfo{ID: id}).First(&chatInfo).Error; err != nil {
		return chatInfo, err
	}

	return chatInfo, nil
}

// Get all ChatInfo
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

// Add ChatCard
func (cs *ChatService) AddChatCard(chatCard model.ChatCard) error {
	tx := global.DB.Begin()

	// 检查ChatInfo是否存在
	if err := tx.Model(&model.ChatInfo{}).Where(&model.ChatInfo{ID: chatCard.ChatInfoID}).First(&model.ChatInfo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := global.DB.Create(&chatCard).Error; err != nil {
		return err
	}

	return nil
}

// Delete ChatCard by id
func (cs *ChatService) DeleteChatCard(id uint) error {
	if err := global.DB.Delete(&model.ChatCard{}, id).Error; err != nil {
		return err
	}

	return nil
}

// Updata ChatCard
func (cs *ChatService) UpdateChatCard(chatCard model.ChatCard) error {
	if err := global.DB.Model(&model.ChatCard{}).Where(&model.ChatCard{ID: chatCard.ID}).Save(&chatCard).Error; err != nil {
		return err
	}

	return nil
}

// Get ChatCard by id
func (cs *ChatService) GetChatCard(id uint) (model.ChatCard, error) {
	chatCard := model.ChatCard{}
	if err := global.DB.Model(&model.ChatCard{}).Where(&model.ChatCard{ID: id}).First(&chatCard).Error; err != nil {
		return chatCard, err
	}

	return chatCard, nil
}

// Get ChatCard by ChatInfo.ID
func (cs *ChatService) GetChatCardByChatInfoID(chatInfoID uint) ([]model.ChatCard, error) {
	chatCards := []model.ChatCard{}
	if err := global.DB.Model(&model.ChatCard{}).Where(&model.ChatCard{ChatInfoID: chatInfoID}).Find(&chatCards).Error; err != nil {
		return nil, err
	}

	return chatCards, nil
}
