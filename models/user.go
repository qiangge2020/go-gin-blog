package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
}

func AddUser(user User) error {
	if err := Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func TakeUser() {

}
