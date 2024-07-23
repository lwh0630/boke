package redis

import (
	"bluebell/models"
	"errors"
	"github.com/redis/go-redis/v9"
)

var (
	ErrSortMethod = errors.New("没有这种排序方法")
)

func GetPostIDListInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	var key string
	if p.Order == models.OrderByTime {
		key = GetRedisKey(KeyPostTime)
	} else if p.Order == models.OrderByScore {
		key = GetRedisKey(KeyPostScore)
	} else {
		return nil, ErrSortMethod
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size
	return rdb.ZRevRange(ctx, key, start, end).Result()
}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) ([]int64, error) {
	//for _, id := range ids {
	//	number := rdb.ZCount(ctx, GetRedisKey(KeyPostVotePrefix+id), "1", "1").Val()
	//	data = append(data, number)
	//}
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		pipeline.ZCount(ctx, GetRedisKey(KeyPostVotePrefix+id), "1", "1")
	}
	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]int64, 0)
	for _, cmd := range cmders {
		number := cmd.(*redis.IntCmd).Val()
		data = append(data, number)
	}

	return data, nil
}
