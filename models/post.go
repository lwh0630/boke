package models

import "time"

type Post struct {
	Id          int64     `json:"id" db:"post_id"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityId int64     `json:"community_id" db:"community_id"`
	Status      int8      `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
}
