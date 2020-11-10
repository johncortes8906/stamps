package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func getUser(userID int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT *
	FROM users
	WHERE id = ?
	`, userID)

	user := &User{}
	err := row.Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Password,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func removeUser(userID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM users WHERE id = ?`, userID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func updateUser(user User) error {
	fmt.Println(user)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if user.ID == nil || *user.ID == 0 {
		return errors.New("user has invalid ID")
	}
	_, err := database.DbConn.ExecContext(ctx, `UPDATE users SET
		fullName=?,
		email=?,
		password=?
		WHERE id=?`,
		user.FullName,
		user.Email,
		user.Password,
		user.ID,
	)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
