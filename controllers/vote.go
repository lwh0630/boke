package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteController(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		var errs validator.ValidationErrors
		if ok := errors.As(err, &errs); !ok {
			zap.L().Error("PostVoteController error", zap.Any("errs", errs))
			ResponseError(c, CodeInvalidParams)
		} else {
			zap.L().Error("validator.ValidateErrors", zap.Any("errs", errs))
			errData := removeTopStruct(errs.Translate(trans))
			ResponseErrorWithMsg(c, CodeInvalidParams, errData)
		}
		return
	}

	uid, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.UserID = uid

	zap.L().Debug("VoteForPostParams", zap.Any("p", p))
	err = logic.VoteForPost(p)
	if err != nil {
		zap.L().Error("logic.VoteForPost error", zap.Any("err", err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseOk(c, nil)
}
