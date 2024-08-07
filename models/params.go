package models

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票参数
type ParamVoteData struct {
	// 从请求中获取当前用户
	UserID    int64 `json:"user_id"`
	PostID    int64 `json:"post_id,string" binding:"required"`       //帖子id
	Direction int8  `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票:1, 反对票:-1, 取消投票:0
}

// ParamPostList 获取帖子列表查询参数
type ParamPostList struct {
	Page  int64  `json:"page,string" form:"page"`
	Size  int64  `json:"size,string" form:"size"`
	Order string `json:"order" form:"order" binding:"oneof=time score"`
}
