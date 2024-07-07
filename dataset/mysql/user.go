package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

// 加密前缀
const secret = "http://www.duxinyu.love"

var (
	ErrUserNotExist    = errors.New("用户不存在")
	ErrUserExist       = errors.New("用户已存在")
	ErrInvalidPassWord = errors.New("用户名或密码错误")
)

// CheckUserExistByUsername 通过用户名，检测用户是否存在
func CheckUserExistByUsername(username string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrUserExist
	}
	return nil
}

// InsertUser 插入用户
func InsertUser(user *models.User) (err error) {
	sqlStr := "insert into user (user_id, username, password) values (?, ?, ?)"
	password := encryptPassword(user.Password)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	return err
}

func Login(p *models.User) (err error) {
	sqlStr := "select user_id, username, password from user where username = ?"
	var user models.User
	err = db.Get(&user, sqlStr, p.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrUserNotExist
	}
	if err != nil {
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(p.Password)
	if password != user.Password {
		return ErrInvalidPassWord
	}
	*p = user
	return
}

// md5加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
