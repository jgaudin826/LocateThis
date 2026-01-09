package dbmodel

import "gorm.io/gorm"

type GroupEntry struct {
	gorm.Model
	Name      string           `json:"name" gorm:"not null"`
	Admin     UserEntry        `json:"admin_id" gorm:"not null;foreignKey:ID;constraint:OnDelete:CASCADE;"`
	Users     []*UserEntry     `gorm:"many2many:group_user_entries;constraint:OnDelete:CASCADE;" json:"users"`
	Locations []*LocationEntry `gorm:"many2many:group_location_entries;constraint:OnDelete:CASCADE;" json:"locations"`
}

type GroupRepository interface {
	Create(entry *GroupEntry) (*GroupEntry, error)
	FindAll() ([]GroupEntry, error)
	FindById(id uint) (*GroupEntry, error)
	FindLocationsForGroup(id uint) ([]LocationEntry, error)
	FindUsersForGroup(id uint) ([]UserEntry, error)
	Update(entry *GroupEntry, id uint) (*GroupEntry, error)
	Delete(id uint) error
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (groupRepository *groupRepository) Create(entry *GroupEntry) (*GroupEntry, error) {
	if err := groupRepository.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (groupRepository *groupRepository) FindAll() ([]GroupEntry, error) {
	var groups []GroupEntry
	if err := groupRepository.db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (groupRepository *groupRepository) FindById(id uint) (*GroupEntry, error) {
	var group GroupEntry
	if err := groupRepository.db.First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (groupRepository *groupRepository) FindLocationsForGroup(id uint) ([]LocationEntry, error) {
	var group GroupEntry
	if err := groupRepository.db.Preload("Locations").First(&group, id).Error; err != nil {
		return nil, err
	}
	locations := make([]LocationEntry, len(group.Locations))
	for i, l := range group.Locations {
		locations[i] = *l
	}
	return locations, nil
}

func (groupRepository *groupRepository) FindUsersForGroup(id uint) ([]UserEntry, error) {
	var group GroupEntry
	if err := groupRepository.db.Preload("Users").First(&group, id).Error; err != nil {
		return nil, err
	}
	users := make([]UserEntry, len(group.Users))
	for i, u := range group.Users {
		users[i] = *u
	}
	return users, nil
}

func (groupRepository *groupRepository) Update(entry *GroupEntry, id uint) (*GroupEntry, error) {
	if err := groupRepository.db.Model(&GroupEntry{}).Where("id = ?", id).Updates(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (groupRepository *groupRepository) Delete(id uint) error {
	if err := groupRepository.db.Delete(&GroupEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}
