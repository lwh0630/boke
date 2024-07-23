package logic

import (
	"bluebell/dataset/redis"
	"bluebell/models"
)

// VoteForPost 投票的具体业务逻辑
// 使用简化版的投票分数
// 投一票加432分, 86400/200 -> 需要200张赞成票可以给帖子续一天.
//
// 投票的几种情况:
//
//  1. Direction = 1
//     之前没有投过票，现在投赞成票
//     之前投反对票，现在改投赞成票
//  2. Direction = 0
//     之前投过赞成票，现在要取消投票
//     之前投过反对票，现在要取消投票
//  3. Direction = -1
//     之前没有投过票，现在投反对票
//     之前投赞成票，现在改投反对票
//
// 投票的限制:
// 每个帖子自发表之日起, 不允许用户再投票
//
//  1. 到期之后将redis中保存的赞成票数以及反对票数存储到mysql中
//  2. 到期之后删除那个 KeyPostVotePrefix
func VoteForPost(p *models.ParamVoteData) error {
	return redis.VoteForPost(p)
}
