package expansion

import (
	"github.com/gin-gonic/gin"
	"niceBackend/common/transform/request"
	"niceBackend/common/transform/response"
	"niceBackend/core/service"
	"niceBackend/utils"
)

func GetStatistic(c *gin.Context) {
	service.GetStatistic()
	response.OkWithMessage("获取成功", c)
}

func PostStatistic(c *gin.Context) {
	var r request.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.AnnouncementVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	service.GetStatistic()
	response.OkWithMessage("写入成功", c)
}