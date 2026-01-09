package dbmodel

import "gorm.io/gorm"

type GroupLocationEntry struct {
	GroupID              uint `gorm:"not null;primaryKey;constraint:OnDelete:CASCADE"`
	LocationID           uint `gorm:"not null;primaryKey;constraint:OnDelete:CASCADE"`
	IsVisibleCoordinates bool `gorm:"not null;default:true"`
}

type GroupLocationRepository interface {
	Create(entry *GroupLocationEntry) (*GroupLocationEntry, error)
	FindAll() ([]GroupLocationEntry, error)
	Update(entry *GroupLocationEntry) (*GroupLocationEntry, error)
	Delete(groupID, locationID uint) error
}
type groupLocationRepository struct {
	db *gorm.DB
}

func NewGroupLocationRepository(db *gorm.DB) GroupLocationRepository {
	return &groupLocationRepository{db: db}
}
func (groupLocationRepository *groupLocationRepository) Create(entry *GroupLocationEntry) (*GroupLocationEntry, error) {
	if err := groupLocationRepository.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (groupLocationRepository *groupLocationRepository) FindAll() ([]GroupLocationEntry, error) {
	var groupLocations []GroupLocationEntry
	if err := groupLocationRepository.db.Find(&groupLocations).Error; err != nil {
		return nil, err
	}
	return groupLocations, nil
}

func (groupLocationRepository *groupLocationRepository) Update(entry *GroupLocationEntry) (*GroupLocationEntry, error) {
	if err := groupLocationRepository.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (groupLocationRepository *groupLocationRepository) Delete(groupID, locationID uint) error {
	return groupLocationRepository.db.Where("group_id = ? AND location_id = ?", groupID, locationID).Delete(&GroupLocationEntry{}).Error
}
