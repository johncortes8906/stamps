package user

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/johncortes8906/stamps/database"
)

var userMap = struct {
	sync.RWMutex
	u map[int]User
}{u: make(map[int]User)}

func getUserList() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT * from users`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	users := make([]User, 0)
	for results.Next() {
		var user User
		results.Scan(
			&user.ID,
			&user.FullName,
			&user.Email,
			&user.Password,
		)
		users = append(users, user)
	}
	return users, nil
}

func insertUser(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO users
	(
		fullName,
		email,
		password
	) VALUES(?, ?, ?)`,
		user.FullName,
		user.Email,
		user.Password,
	)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	userID, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return int(userID), nil
}
