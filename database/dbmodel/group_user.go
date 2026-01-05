package dbmodel

import "gorm.io/gorm"

type GroupUserEntry struct {
	UserID  uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	GroupID uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
}

type UserGroupRepository interface {
	Create(entry *GroupUserEntry) (*GroupUserEntry, error)
	FindAll() ([]GroupUserEntry, error)
	Delete(userID, groupID uint) error
}

type userGroupRepository struct {
	db *gorm.DB
}

func NewUserGroupRepository(db *gorm.DB) UserGroupRepository {
	return &userGroupRepository{db: db}
}

func (userGroupRepository *userGroupRepository) Create(entry *GroupUserEntry) (*GroupUserEntry, error) {
	if err := userGroupRepository.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (userGroupRepository *userGroupRepository) FindAll() ([]GroupUserEntry, error) {
	var userGroups []GroupUserEntry
	if err := userGroupRepository.db.Find(&userGroups).Error; err != nil {
		return nil, err
	}
	return userGroups, nil
}

func (userGroupRepository *userGroupRepository) Delete(userID, groupID uint) error {
	return userGroupRepository.db.Where("user_id = ? AND group_id = ?", userID, groupID).Delete(&GroupUserEntry{}).Error
}
