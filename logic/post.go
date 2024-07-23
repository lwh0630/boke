package logic

import (
	"bluebell/dataset/mysql"
	"bluebell/dataset/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) error {
	// 1. 生成post id
	p.Id = snowflake.GenID()
	// 2. 保存到mysql
	if err := mysql.CreatePost(p); err != nil {
		return err
	}
	// 3. 保存到redis
	if err := redis.CreatePost(p.Id); err != nil {
		return err
	}
	return nil
}

// GetPostById 根据帖子id，查询帖子详情
func GetPostById(id int64) (*models.ApiPostDetail, error) {
	// 获取帖子详情
	post, err := mysql.GetPostById(id)
	if err != nil {
		return nil, err
	}
	apiPostDetail := &models.ApiPostDetail{
		Post: post,
	}

	// 获取作者昵称
	user, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		return nil, err
	}
	apiPostDetail.AuthorName = user.Username

	// 获取社区信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityId)
	if err != nil {
		return nil, err
	}
	apiPostDetail.CommunityDetail = community

	return apiPostDetail, nil
}

// GetPostList 返回帖子列表
func GetPostList(pageNum int64, pageSize int64) ([]*models.ApiPostDetail, error) {
	// 获取帖子
	posts, err := mysql.GetPostList((pageNum-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	apiPostList := make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 获取作者昵称
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			return nil, err
		}
		// 获取社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityId)
		if err != nil {
			return nil, err
		}

		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		apiPostList = append(apiPostList, postDetail)
	}
	return apiPostList, nil
}
