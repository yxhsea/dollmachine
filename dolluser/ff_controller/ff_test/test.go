package ff_test

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

// @Summary 测试专用
// @tags test
// @Produce  json
// @Param Token header string true "Token"
// @Param test query string true "Test"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /test/info [get]
func GetTest(ctx *gin.Context) {
	fmt.Println("test", ctx.Query("test"))
	fmt.Println("token", ctx.Request.Header.Get("Token"))

	urlPath := ctx.Request.URL.Path
	fmt.Printf("urlPath %v \n", urlPath)

	ctx.JSON(200, gin.H{
		"message": "test",
	})
}