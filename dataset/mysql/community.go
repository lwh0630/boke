package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
)

var (
	Err = errors.New("text")
)

func GetCommunityList() (communityList []models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
	}
}
