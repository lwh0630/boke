package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	fmt.Println(communityList)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("there is no community", zap.Error(err))
			err = nil
			communityList = []models.Community{}
		} else {
			err = ErrSelectCommunityList
		}
	}
	return communityList, err
}

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = &models.CommunityDetail{}
	sqlStr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
	err = db.Get(community, sqlStr, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrInvalidCommunityId
		}
	}
	return community, err
}
