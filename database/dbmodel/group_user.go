package dbmodel

import "gorm.io/gorm"

type GroupUserEntry struct {
	UserID  uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	GroupID uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
}

type GroupUserRepository interface {
	Create(entry *GroupUserEntry) (*GroupUserEntry, error)
	FindAll() ([]GroupUserEntry, error)
	Delete(userID, groupID uint) error
}

type groupUserRepository struct {
	db *gorm.DB
}

func NewGroupUserRepository(db *gorm.DB) GroupUserRepository {
	return &groupUserRepository{db: db}
}

func (groupUserRepository *groupUserRepository) Create(entry *GroupUserEntry) (*GroupUserEntry, error) {
	if err := groupUserRepository.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (groupUserRepository *groupUserRepository) FindAll() ([]GroupUserEntry, error) {
	var userGroups []GroupUserEntry
	if err := groupUserRepository.db.Find(&userGroups).Error; err != nil {
		return nil, err
	}
	return userGroups, nil
}

func (groupUserRepository *groupUserRepository) Delete(userID, groupID uint) error {
	return groupUserRepository.db.Where("user_id = ? AND group_id = ?", userID, groupID).Delete(&GroupUserEntry{}).Error
}
