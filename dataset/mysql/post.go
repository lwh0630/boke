package mysql

import (
	"bluebell/models"
	"errors"
	"github.com/jmoiron/sqlx"
	"strings"
)

var NoData = errors.New("没有数据")

// CreatePost 创建帖子
func CreatePost(p *models.Post) error {
	sqlStr := "insert into post (post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.Id, p.Title, p.Content, p.AuthorId, p.CommunityId)
	if err != nil {
		return err
	}
	return nil
}

// GetPostById 根据id查询单个帖子的数据
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
	sqlStr := "select post_id, title, content, author_id, community_id, create_time, update_time from post order by post.create_time DESC limit ? offset ?"
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

// GetPostByIndexes 根据id列表查询
func GetPostByIndexes(indexes []string) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time, update_time 
				from post 
				where post_id in (?)
				order by FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, indexes, strings.Join(indexes, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&posts, query, args...)
	return posts, err
}
