package option


import (
	"github.com/gin-gonic/gin"
	"niceBackend/common/transform/response"
)

func GetOption(c *gin.Context){
	response.OkWithMessage("ok", c)
	return
}
