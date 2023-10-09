package models

import (
	"github.com/azkazkazka/task-todo/db"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FetchUsers() (Response, error) {
	var res Response
	var users []User
	con := db.CreateCon()

	if err := con.Find(&users).Error; err != nil {
		return res, err
	}

	res.Message = "Successfully fetched users"
	res.Data = users

	return res, nil
}

func CreateUser(user *User) (Response, error) {
	var res Response
	con := db.CreateCon()

	if err := con.Create(user).Error; err != nil {
		return res, err
	}

	res.Message = "Successfully created user"
	return res, nil
}

func DeleteUser(userID uint) (Response, error) {
	var res Response
	con := db.CreateCon()
	existingUser := &User{}

	if err := con.First(existingUser, userID).Error; err != nil {
		return res, err
	}

	if err := con.Delete(existingUser).Error; err != nil {
		return res, err
	}

	res.Message = "Successfully deleted user"
	return res, nil
}

func UpdateUser(user *User) (Response, error) {
	var res Response
	con := db.CreateCon()

	existingUser := &User{}
	if err := con.First(existingUser, user.ID).Error; err != nil {
		return res, err
	}

	if err := con.Model(&User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return res, err
	}

	res.Message = "Successfully updated user"
	return res, nil
}
