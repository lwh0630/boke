package redis

import (
	"bluebell/models"
	"errors"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"math"
	"strconv"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 60 * 60
	scorePerVote     = 432
)

var (
	ctx            = context.Background()
	ErrVoteTimeout = errors.New("超出投票时间")
)

func CreatePost(postID int64) error {
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(ctx, GetRedisKey(KeyPostTime), redis.Z{Score: float64(time.Now().Unix()), Member: postID})
	pipeline.ZAdd(ctx, GetRedisKey(KeyPostScore), redis.Z{Score: .0, Member: postID})

	_, err := pipeline.Exec(ctx)
	return err
}

func VoteForPost(p *models.ParamVoteData) (err error) {
	postID := strconv.FormatInt(p.PostID, 10)
	userID := strconv.FormatInt(p.UserID, 10)
	value := float64(p.Direction)
	// 1. 判断投票的限制
	// 从redis中获取帖子是否过期
	postTime := rdb.ZScore(ctx, GetRedisKey(KeyPostTime), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeout
	}
	// 2. 更新帖子的分数
	// 先查当前用户给当前帖子的投票记录
	pipeline := rdb.TxPipeline()
	oldValue := rdb.ZScore(ctx, GetRedisKey(KeyPostVotePrefix+postID), userID).Val()
	diff := math.Abs(oldValue - value)
	if value > oldValue {
		pipeline.ZIncrBy(ctx, GetRedisKey(KeyPostScore), diff*scorePerVote, postID)
	} else {
		pipeline.ZIncrBy(ctx, GetRedisKey(KeyPostScore), -1*diff*scorePerVote, postID)
	}
	// 3. 记录用户为帖子投票的记录
	if value == 0 {
		pipeline.ZRem(ctx, GetRedisKey(KeyPostVotePrefix+postID), userID)
	} else {
		pipeline.ZAdd(ctx, GetRedisKey(KeyPostVotePrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err = pipeline.Exec(ctx)
	return err
}
