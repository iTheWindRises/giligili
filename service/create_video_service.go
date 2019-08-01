package service

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateVideoService 视频投稿服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=1,max=20"`
	Info  string `form:"info" json:"info" binding:"required,min=0,max=256"`
}

// Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "ok",
		Data: serializer.BuildVideo(video),
	}

}
