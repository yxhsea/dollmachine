package ff_v1test

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summary 测试专用
// @tags test
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param test formData string true "Test"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /test/info [put]
func GetTest(ctx *gin.Context) {
	fmt.Println(ctx.PostForm("test"))
	ctx.JSON(200, gin.H{
		"message": "test",
	})
}
