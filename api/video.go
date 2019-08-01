package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// 添加视频,视频投稿
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}

	if err := c.ShouldBind(&service); err == nil {
		result := service.Create()
		c.JSON(200, result)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//视频详情接口
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}

	if err := c.ShouldBind(&service); err == nil {
		result := service.Show(c.Param("id"))
		c.JSON(200, result)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//视频详情列表接口
func ListVideo(c *gin.Context) {
	service := service.ShowVideoService{}

	if err := c.ShouldBind(&service); err == nil {
		result := service.ShowList()
		c.JSON(200, result)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 视频更新接口
func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}

	if err := c.ShouldBind(&service); err == nil {
		result := service.Update(c.Param("id"))
		c.JSON(200, result)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 视频删除接口
func DeleteVideo(c *gin.Context) {
	service := service.ShowVideoService{}

	result := service.Delete(c.Param("id"))
	c.JSON(200, result)

}
