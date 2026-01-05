package dbmodel

import "gorm.io/gorm"

type GroupEntry struct {
	gorm.Model
	Name      string           `json:"name"`
	Users     []*UserEntry     `gorm:"many2many:group_users;constraint:OnDelete:CASCADE;" json:"users"`
	Locations []*LocationEntry `gorm:"many2many:group_locations;" json:"locations"`
}

type GroupRepository interface {
	Create(entry *GroupEntry) (*GroupEntry, error)
	FindById(id uint) (*GroupEntry, error)
	Update(entry *GroupEntry) (*GroupEntry, error)
	Delete(id uint) error
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) Create(entry *GroupEntry) (*GroupEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *groupRepository) FindById(id uint) (*GroupEntry, error) {
	var group GroupEntry
	if err := r.db.First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *groupRepository) Update(entry *GroupEntry) (*GroupEntry, error) {
	if err := r.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *groupRepository) Delete(id uint) error {
	if err := r.db.Delete(&GroupEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}
