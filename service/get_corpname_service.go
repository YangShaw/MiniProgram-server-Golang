package service

import (
	"Miniprogram-server-Golang/model"
	"Miniprogram-server-Golang/serializer"

	"github.com/gin-gonic/gin"
)

// GetCorpService 管理用户企业身份服务
type GetCorpService struct {
	Uid    string `form:"uid" json:"uid"`
	Token  string `form:"token" json:"token"`
	Corpid string `form:"corpid" json:"corpid"`
}

// GetCorp 获取用户企业信息
func (service *GetCorpService) GetCorp(c *gin.Context) serializer.Response {

	if !model.CheckToken(service.Uid, service.Token) {
		return serializer.ParamErr("token验证错误", nil)
	}

	var corp model.Corp
	if err := model.DB.Where(&model.Corp{Corpid: service.Corpid}).First(&corp).Error; err != nil {
		return serializer.Err(10006, "获取企业信息失败", nil)
	}

	return serializer.BuildCorpResponse(0, corp)
}
