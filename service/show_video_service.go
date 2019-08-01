package service

import (
	"giligili/model"
	"giligili/serializer"
	"strconv"
)

// ShowVideoService 视频服务
type ShowVideoService struct{}

// Show 视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video

	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "ok",
		Data: serializer.BuildVideo(video),
	}

}

// Show 视频
func (service *ShowVideoService) ShowList() serializer.Response {
	var video []model.Video

	err := model.DB.Limit(2).Find(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  5000,
			Msg:   "数据库错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "ok",
		Data: serializer.BuildVideoList(video),
	}

}

// Delete 视频
func (service *ShowVideoService) Delete(id string) serializer.Response {

	var video model.Video

	uid, _ := strconv.Atoi(id)

	video.ID = uint(uid)
	err := model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  5000,
			Msg:   "数据库错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 0,
		Msg:  "ok",
	}

}
