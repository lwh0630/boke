package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数以及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("create post with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 从context中拿到用户id
	UseID, err := GetCurrentUserID(c)
	if err != nil {
		fmt.Println(err)
		zap.L().Error("get current user id failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	println("123")
	p.AuthorId = UseID
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseOk(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数
	postIdStr := c.Param("id")
	if postIdStr == "" {
		zap.L().Error("get post id failed", zap.String("param", c.Param("id")))
		ResponseError(c, CodeInvalidParams)
		return
	}
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		zap.L().Error("convert post id failed", zap.String("param", c.Param("id")))
		ResponseError(c, CodeInvalidParams)
	}
	// 2. 根据id取出帖子数据
	post, err := logic.GetPostById(postId)
	if err != nil {
		zap.L().Error("logic.GetPostById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseOk(c, post)
}

// GetPostListHandler 根据参数获取帖子列表
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	pageNumStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("size", "10")

	pageNum, _ := strconv.ParseInt(pageNumStr, 10, 64)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 64)
	zap.L().Info("page", zap.Int64("pageNum", pageNum), zap.Int64("pageSize", pageSize))

	if pageNum <= 0 || pageSize <= 0 {
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 获取数据
	data, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseOk(c, data)
}

// GetPostListHandlerV2 根据参数获取帖子列表V2，多加排序功能
func GetPostListHandlerV2(c *gin.Context) {
	// 获取分页参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderByTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("get post list with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	zap.L().Info("page", zap.Int64("pageNum", p.Page), zap.Int64("pageSize", p.Size), zap.String("order", p.Order))

	if p.Page <= 0 || p.Size <= 0 {
		zap.L().Error("get post list with invalid param", zap.Any("param", p))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 获取数据
	data, err := logic.GetPostListV2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseOk(c, data)
}
