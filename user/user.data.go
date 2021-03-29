package user

import (
	"errors"
	"fmt"

	"github.com/johncortes8906/stamps/database"
)

var (
	user  User
	users []User
)

func getUserList() ([]User, error) {
	database.Db.Find(&users)
	return users, nil
}

func insertUser(userRequest User) (*int, error) {
	database.Db.Create(&userRequest)
	return userRequest.ID, nil
}

func getUser(userID int) (User, error) {
	database.Db.First(&user, userID)
	return user, nil
}

func removeUser(userID int) error {
	database.Db.Delete(&user, userID)
	return nil
}

func updateUser(userRequest User) error {
	fmt.Println(*userRequest.ID)
	if userRequest.ID == nil || *userRequest.ID == 0 {
		return errors.New("user has invalid ID")
	}
	database.Db.First(&userRequest)
	database.Db.Save(&userRequest)
	return nil
}
