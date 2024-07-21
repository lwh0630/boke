package mysql

import (
	"bluebell/models"
	"errors"
)

var NoData = errors.New("没有数据")

func CreatePost(p *models.Post) error {
	sqlStr := "insert into post (post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.Id, p.Title, p.Content, p.AuthorId, p.CommunityId)
	if err != nil {
		return err
	}
	return nil
}

func GetPostById(id int64) (*models.Post, error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time, update_time from post where post_id = ?"
	post := &models.Post{}

	err := db.Get(post, sqlStr, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// GetPostList 获取帖子列表
// offset 索引起点, limit 索引大小
func GetPostList(offset int64, limit int64) ([]*models.Post, error) {
	sqlStr := "select post_id, title, content, author_id, community_id, create_time, update_time from post limit ? offset ?"
	var posts []*models.Post
	err := db.Select(&posts, sqlStr, limit, offset)
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return nil, NoData
	}
	return posts, nil
}
