package user

import "gorm.io/gorm"

type UserRepo struct {
	Db *gorm.DB
}

func (userRepo *UserRepo) Create(user User) error {
	return userRepo.Db.Create(&user).Error
}
