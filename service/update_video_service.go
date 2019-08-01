package service

import (
	"giligili/model"
	"giligili/serializer"
)

// UpdateVideoService 视频投稿服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"min=2,max=20"`
	Info  string `form:"info" json:"info" binding:"min=0,max=256"`
}

// Update 视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	err := model.DB.Model(&video).Where("id=?", id).Update(video).Error
	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "视频更新失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "ok",
		Data: serializer.BuildVideo(video),
	}

}
