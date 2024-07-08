package logic

import (
	"bluebell/dataset/mysql"
	"bluebell/models"
)

func GetCommunityList() (communityList []models.Community, err error) {
	// 查询数据库，查找全部的community
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
