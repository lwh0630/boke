package controllers

import (
	"bluebell/dataset/mysql"
	"bluebell/logic"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CommunityHandler 社区相关
func CommunityHandler(c *gin.Context) {
	// 查询到全部的社区 (community_id, community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseOk(c, data)
}

// CommunityDetailHandler 根据社区ID查询社区详情
func CommunityDetailHandler(c *gin.Context) {
	// 1.获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("id 类型转换错误", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail(id) failed", zap.Error(err))
		if errors.Is(err, mysql.ErrInvalidCommunityId) {
			ResponseError(c, CodeInvalidParams)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseOk(c, data)
}
