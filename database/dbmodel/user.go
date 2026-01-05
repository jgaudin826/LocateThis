package dbmodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Pseudo   string `json:"pseudo"`
}

type UserRepository interface {
	Create(entry *User) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(entry *User) (*User, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindById(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindLocationsForUser(id uint) ([]Location, error) {
	var locations []Location
	if err := r.db.Where("id_user = ?", id).Find(&locations).Error; err != nil {
		return nil, err
	}
	return locations, nil
}

func (r *userRepository) Delete(id uint) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(entry *User) (*User, error) {
	if err := r.db.Save(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}
