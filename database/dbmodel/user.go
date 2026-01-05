package dbmodel

import "gorm.io/gorm"

type UserEntry struct {
	gorm.Model
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Pseudo   string        `json:"pseudo"`
	Groups   []*GroupEntry `gorm:"many2many:group_users;constraint:OnDelete:CASCADE;" json:"groups"`
}

type UserRepository interface {
	Create(entry *UserEntry) (*UserEntry, error)
	FindAll() ([]UserEntry, error)
	FindById(id uint) (*UserEntry, error)
	FindByEmail(email string) (*UserEntry, error)
	FindLocationsForUser(id uint) ([]LocationEntry, error)
	FindGroupsForUser(id uint) ([]GroupEntry, error)
	Update(entry *UserEntry) (*UserEntry, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(entry *UserEntry) (*UserEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *userRepository) FindAll() ([]UserEntry, error) {
	var users []UserEntry
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindById(id uint) (*UserEntry, error) {
	var user UserEntry
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*UserEntry, error) {
	var user UserEntry
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindLocationsForUser(id uint) ([]LocationEntry, error) {
	var locations []LocationEntry
	if err := r.db.Where("id_user = ?", id).Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func (r *userRepository) FindGroupsForUser(id uint) ([]GroupEntry, error) {
	var groups []GroupEntry
	if err := r.db.Where("id_user = ?", id).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *userRepository) Update(entry *UserEntry) (*UserEntry, error) {
	if err := r.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *userRepository) Delete(id uint) error {
	if err := r.db.Delete(&UserEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}
