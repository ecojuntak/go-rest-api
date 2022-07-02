package user

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

type Repository interface {
	GetTotalUser() (total int64, err error)
	Create(user User) (err error)
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) GetTotalUser() (total int64, err error) {
	err = u.db.Model(&User{}).Count(&total).Error
	return
}

func (u userRepository) Create(user User) (err error) {
	return u.db.Create(&user).Error
}
