package dbmodel

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	ID_User   User     `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;" json:"id_user"`
	Latitude  string   `json:"latitude"`
	Longitude string   `json:"longitude"`
	Name      string   `json:"name"`
	Groups    []*Group `gorm:"many2many:group_locations;constraint:OnDelete:CASCADE;" json:"groups"`
}

type LocationRepository interface {
	Create(entry *Location) (*Location, error)
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (locationRepository *locationRepository) Create(entry *Location) (*Location, error) {
	if err := locationRepository.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (locationRepository *locationRepository) FindAll() ([]Location, error) {
	var locations []Location
	if err := locationRepository.db.Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func (locationRepository *locationRepository) FindById(id uint) (*Location, error) {
	var location Location
	if err := locationRepository.db.First(&location, id).Error; err != nil {
		return nil, err
	}
	return &location, nil
}

func (locationRepository *locationRepository) Delete(id uint) error {
	if err := locationRepository.db.Delete(&Location{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (locationRepository *locationRepository) Update(entry *Location) (*Location, error) {
	if err := locationRepository.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}
