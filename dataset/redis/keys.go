package redis

// redis key 使用命名空间的方式，方便查询和拆分
// ZSet: 有序集合
const (
	KeyPrefix         = "bluebell:"
	KeyPostTime       = "post:time"  // ZSet;帖子发帖时间
	KeyPostScore      = "post:score" // ZSet;帖子及投票的分数
	KeyPostVotePrefix = "post:vote:" // ZSet;记录用户以及投票类型;参数是post id
)

// GetRedisKey 返回加上命名空间的索引
func GetRedisKey(key string) string {
	return KeyPrefix + key
}
