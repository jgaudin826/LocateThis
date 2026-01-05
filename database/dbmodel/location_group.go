package dbmodel

import "gorm.io/gorm"

type GroupLocation struct {
	GroupID              uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	LocationID           uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	IsVisibleCoordinates bool `gorm:"default:true"`
}

type LocationGroupRepository interface {
	Create(entry *GroupLocation) (*GroupLocation, error)
	Delete(groupID, locationID uint) error
}
type locationGroupRepository struct {
	db *gorm.DB
}

func NewLocationGroupRepository(db *gorm.DB) LocationGroupRepository {
	return &locationGroupRepository{db: db}
}
func (r *locationGroupRepository) Create(entry *GroupLocation) (*GroupLocation, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *locationGroupRepository) Update(entry *GroupLocation) (*GroupLocation, error) {
	if err := r.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *locationGroupRepository) Delete(groupID, locationID uint) error {
	return r.db.Where("group_id = ? AND location_id = ?", groupID, locationID).Delete(&GroupLocation{}).Error
}
