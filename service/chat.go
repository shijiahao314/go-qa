package service

import (
	"fmt"

	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"gorm.io/gorm"
)

type ChatService struct{}

// Check userId是否与ChatInfo.UserID匹配
func (cs *ChatService) CheckUser(userID uint64, id uint) error {
	var chatInfo model.ChatInfo
	if err := global.DB.Model(&model.ChatInfo{}).Where(&model.ChatInfo{ID: id}).Take(&chatInfo).Error; err != nil {
		return err
	}
	if userID != chatInfo.UserID {
		return fmt.Errorf("user id [%d] does not match chat info user id [%d]", userID, chatInfo.UserID)
	}
	return nil
}

// Add ChatInfo
func (cs *ChatService) AddChatInfo(chatInfo *model.ChatInfo) error {
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
	tx := global.DB.Begin()
	oldChatInfo := model.ChatInfo{}
	if err := tx.Model(&model.ChatInfo{}).Where("id = ?", chatInfo.ID).Take(&oldChatInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	// set new data
	oldChatInfo.Title = chatInfo.Title

	if err := global.DB.Model(&model.ChatInfo{}).Where("id = ?", oldChatInfo.ID).Save(&oldChatInfo).Error; err != nil {
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
	if err := tx.Model(&model.ChatInfo{}).Where(&model.ChatInfo{UserID: userId}).Order("updated_at desc").Find(&chatInfos).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return chatInfos, nil
}

// Add ChatCard
func (cs *ChatService) AddChatCard(chatCard *model.ChatCard) error {
	tx := global.DB.Begin()

	// 检查ChatInfo是否存在
	if err := tx.Model(&model.ChatInfo{}).Where(&model.ChatInfo{ID: chatCard.ChatInfoID}).First(&model.ChatInfo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&chatCard).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.ChatInfo{}).Where("id = ?", chatCard.ChatInfoID).Update("num", gorm.Expr("num + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

// Delete ChatCard by id
func (cs *ChatService) DeleteChatCard(id uint) error {
	tx := global.DB.Begin()

	var chatInfoID uint
	if err := tx.Model(&model.ChatCard{}).Where(&model.ChatCard{ID: id}).Pluck("chat_info_id", &chatInfoID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.ChatCard{}).Delete(&model.ChatCard{ID: id}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.ChatInfo{}).Where(&model.ChatInfo{ID: chatInfoID}).Update("num", gorm.Expr("num - 1")).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

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

// Get ChatCards by ChatID
func (cs *ChatService) GetChatCards(id uint) ([]model.ChatCard, error) {
	chatCards := []model.ChatCard{}
	if err := global.DB.Model(&model.ChatCard{}).Where(&model.ChatCard{ChatInfoID: id}).Find(&chatCards).Error; err != nil {
		return nil, err
	}
	return chatCards, nil
}

// Get ChatCard by ChatInfo.ID
func (cs *ChatService) GetChatCardByChatInfoID(chatInfoID uint) ([]model.ChatCard, error) {
	chatCards := []model.ChatCard{}
	if err := global.DB.Model(&model.ChatCard{}).Where(&model.ChatCard{ChatInfoID: chatInfoID}).Find(&chatCards).Error; err != nil {
		return nil, err
	}

	return chatCards, nil
}
