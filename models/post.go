package models

import "time"

type Post struct {
	Id          int64     `json:"id,string" db:"post_id"`
	AuthorId    int64     `json:"author_id,string" db:"author_id"`
	CommunityId int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int8      `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
}

// ApiPostDetail 帖子详情接口结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"`
	*Post                               //嵌入帖子结构体
	*CommunityDetail `json:"community"` //嵌入社区信息
}
