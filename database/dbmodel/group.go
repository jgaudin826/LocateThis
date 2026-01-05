package dbmodel

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name      string      `json:"name"`
	Users     []*User     `gorm:"many2many:group_users;constraint:OnDelete:CASCADE;" json:"users"`
	Locations []*Location `gorm:"many2many:group_locations;" json:"locations"`
}

type GroupRepository interface {
	Create(entry *Group) (*Group, error)
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) Create(entry *Group) (*Group, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *groupRepository) FindById(id uint) (*Group, error) {
	var group Group
	if err := r.db.First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *groupRepository) Delete(id uint) error {
	if err := r.db.Delete(&Group{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *groupRepository) Update(entry *Group) (*Group, error) {
	if err := r.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}
