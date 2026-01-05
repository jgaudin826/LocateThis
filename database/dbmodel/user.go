package dbmodel

import "gorm.io/gorm"

type UserEntry struct {
	gorm.Model
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Username string        `json:"username"`
	Groups   []*GroupEntry `gorm:"many2many:group_users;constraint:OnDelete:CASCADE;" json:"groups"`
}

type UserRepository interface {
	Create(entry *UserEntry) (*UserEntry, error)
	FindAll() ([]UserEntry, error)
	FindById(id uint) (*UserEntry, error)
	FindByEmail(email string) (*UserEntry, error)
	FindByUsername(username string) (*UserEntry, error)
	FindLocationsForUser(id uint) ([]LocationEntry, error)
	FindGroupsForUser(id uint) ([]GroupEntry, error)
	Update(entry *UserEntry, id uint) (*UserEntry, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) Create(entry *UserEntry) (*UserEntry, error) {
	if err := userRepository.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (userRepository *userRepository) FindAll() ([]UserEntry, error) {
	var users []UserEntry
	if err := userRepository.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepository *userRepository) FindById(id uint) (*UserEntry, error) {
	var user UserEntry
	if err := userRepository.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) FindByEmail(email string) (*UserEntry, error) {
	var user UserEntry
	if err := userRepository.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) FindByUsername(username string) (*UserEntry, error) {
	var user UserEntry
	if err := userRepository.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) FindLocationsForUser(id uint) ([]LocationEntry, error) {
	var locations []LocationEntry
	if err := userRepository.db.Where("id_user = ?", id).Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func (userRepository *userRepository) FindGroupsForUser(id uint) ([]GroupEntry, error) {
	var user UserEntry
	if err := userRepository.db.Preload("Groups").First(&user, id).Error; err != nil {
		return nil, err
	}
	groups := make([]GroupEntry, len(user.Groups))
	for i, g := range user.Groups {
		groups[i] = *g
	}
	return groups, nil
}

func (userRepository *userRepository) Update(entry *UserEntry, id uint) (*UserEntry, error) {
	if err := userRepository.db.Model(&UserEntry{}).Where("id = ?", id).Updates(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (userRepository *userRepository) Delete(id uint) error {
	if err := userRepository.db.Delete(&UserEntry{}, id).Error; err != nil {
		return err
	}
	return nil
}
