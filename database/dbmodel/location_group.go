package dbmodel

import "gorm.io/gorm"

type GroupLocationEntry struct {
	GroupID              uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	LocationID           uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	IsVisibleCoordinates bool `gorm:"default:true"`
}

type LocationGroupRepository interface {
	Create(entry *GroupLocationEntry) (*GroupLocationEntry, error)
	Update(entry *GroupLocationEntry) (*GroupLocationEntry, error)
	Delete(groupID, locationID uint) error
}
type locationGroupRepository struct {
	db *gorm.DB
}

func NewLocationGroupRepository(db *gorm.DB) LocationGroupRepository {
	return &locationGroupRepository{db: db}
}
func (r *locationGroupRepository) Create(entry *GroupLocationEntry) (*GroupLocationEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *locationGroupRepository) Update(entry *GroupLocationEntry) (*GroupLocationEntry, error) {
	if err := r.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *locationGroupRepository) Delete(groupID, locationID uint) error {
	return r.db.Where("group_id = ? AND location_id = ?", groupID, locationID).Delete(&GroupLocationEntry{}).Error
}
