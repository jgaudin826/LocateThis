package dbmodel

import "gorm.io/gorm"

type userGroup struct {
	UserID  uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	GroupID uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
}

type UserGroupRepository interface {
	Create(entry *userGroup) (*userGroup, error)
	FindAll() ([]userGroup, error)
	Delete(userID, groupID uint) error
}

type userGroupRepository struct {
	db *gorm.DB
}

func NewUserGroupRepository(db *gorm.DB) UserGroupRepository {
	return &userGroupRepository{db: db}
}

func (r *userGroupRepository) Create(entry *userGroup) (*userGroup, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *userGroupRepository) FindAll() ([]userGroup, error) {
	var userGroups []userGroup
	if err := r.db.Find(&userGroups).Error; err != nil {
		return nil, err
	}
	return userGroups, nil
}

func (r *userGroupRepository) Delete(userID, groupID uint) error {
	return r.db.Where("user_id = ? AND group_id = ?", userID, groupID).Delete(&userGroup{}).Error
}
