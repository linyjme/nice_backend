package category

import (
	"github.com/gin-gonic/gin"
	"niceBackend/common/transform/response"
)

func GetCategory(c *gin.Context){

	var result map[string]interface{}
	result = make(map[string]interface{})
	var data [1]int
	result["data"] = data
	response.OkWithData(result, c)
	return
}

