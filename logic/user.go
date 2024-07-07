package logic

import (
	"bluebell/dataset/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户存不存在
	err = mysql.CheckUserExistByUsername(p.Username)
	if err != nil {
		return err
	}
	// 2. 生成UID
	userID := snowflake.GenID()
	// 构造一个实例
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
		Email:    "",
		Gender:   "",
	}
	// 3. 保存进数据库
	err = mysql.InsertUser(&user)
	return err
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 生成JWT token
	return jwt.GenToken(user.UserID, user.Username)
}
