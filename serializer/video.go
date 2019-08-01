package serializer

import (
	"giligili/model"
)

// User 用户序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildVideo 序列化视频
func BuildVideoList(list []model.Video) []Video {
	var video []Video

	for _, item := range list {
		v := Video{
			ID:        item.ID,
			Title:     item.Title,
			Info:      item.Info,
			CreatedAt: item.CreatedAt.Unix(),
		}
		video = append(video, v)
	}
	return video
}
