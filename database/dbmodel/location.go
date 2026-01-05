package dbmodel

import "gorm.io/gorm"

type LocationEntry struct {
	gorm.Model
	User      UserEntry     `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;" json:"id_user"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Name      string        `json:"name"`
	Groups    []*GroupEntry `gorm:"many2many:group_locations;constraint:OnDelete:CASCADE;" json:"groups"`
}

type LocationRepository interface {
	Create(entry *LocationEntry) (*LocationEntry, error)
	FindAll() ([]LocationEntry, error)
	FindById(id uint) (*LocationEntry, error)
	FindGroupsForLocation(id uint) ([]GroupEntry, error)
	Update(entry *LocationEntry, id uint) (*LocationEntry, error)
	Delete(id uint) error
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (locationRepository *locationRepository) Create(entry *LocationEntry) (*LocationEntry, error) {
	if err := locationRepository.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (locationRepository *locationRepository) FindAll() ([]LocationEntry, error) {
	var locations []LocationEntry
	if err := locationRepository.db.Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func (locationRepository *locationRepository) FindById(id uint) (*LocationEntry, error) {
	var location LocationEntry
	if err := locationRepository.db.First(&location, id).Error; err != nil {
		return nil, err
	}
	return &location, nil
}

func (locationRepository *locationRepository) FindGroupsForLocation(id uint) ([]GroupEntry, error) {
	var location LocationEntry
	if err := locationRepository.db.Preload("Groups").First(&location, id).Error; err != nil {
		return nil, err
	}
	groups := make([]GroupEntry, len(location.Groups))
	for i, g := range location.Groups {
		groups[i] = *g
	}
	return groups, nil
}

func (locationRepository *locationRepository) Update(entry *LocationEntry, id uint) (*LocationEntry, error) {
	if err := locationRepository.db.Model(&LocationEntry{}).Where("id = ?", id).Updates(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (locationRepository *locationRepository) Delete(id uint) error {
	if err := locationRepository.db.Delete(&LocationEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}
