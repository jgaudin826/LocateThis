package dbmodel

import "gorm.io/gorm"

type GroupLocationEntry struct {
	GroupID              uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	LocationID           uint `gorm:"primaryKey;constraint:OnDelete:CASCADE"`
	IsVisibleCoordinates bool `gorm:"default:true"`
}

type LocationGroupRepository interface {
	Create(entry *GroupLocationEntry) (*GroupLocationEntry, error)
	FindAll() (*GroupLocationEntry, error)
	Update(entry *GroupLocationEntry) (*GroupLocationEntry, error)
	Delete(groupID, locationID uint) error
}
type locationGroupRepository struct {
	db *gorm.DB
}

func NewLocationGroupRepository(db *gorm.DB) LocationGroupRepository {
	return &locationGroupRepository{db: db}
}
func (locationGroupRepository *locationGroupRepository) Create(entry *GroupLocationEntry) (*GroupLocationEntry, error) {
	if err := locationGroupRepository.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (locationGroupRepository *locationGroupRepository) FindAll() (*GroupLocationEntry, error) {
	var groupLocation GroupLocationEntry
	if err := locationGroupRepository.db.First(&groupLocation).Error; err != nil {
		return nil, err
	}
	return &groupLocation, nil
}

func (locationGroupRepository *locationGroupRepository) Update(entry *GroupLocationEntry) (*GroupLocationEntry, error) {
	if err := locationGroupRepository.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (locationGroupRepository *locationGroupRepository) Delete(groupID, locationID uint) error {
	return locationGroupRepository.db.Where("group_id = ? AND location_id = ?", groupID, locationID).Delete(&GroupLocationEntry{}).Error
}
