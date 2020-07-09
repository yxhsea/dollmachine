package ff_v1upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_common/ff_random"
	"dollmachine/dollmerchant/ff_common/ff_upload"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"
)

// @Summary 上传图片
// @tags Upload
// @Accept  multipart/form-data
// @Produce  json
// @Param Token header string true "Token"
// @Param file formData file true "图片"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload/img [post]
func UploadImg(ctx *gin.Context) {
	var code int
	file, _ := ctx.FormFile("file")
	imgUrl, err := UploadImgQiuNiu(file)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "图片上传失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "图片上传成功", "data": map[string]string{"img_url": imgUrl}})
	return
}

//上传图片到七牛
func UploadImgQiuNiu(file *multipart.FileHeader) (string, error) {
	FileSuffix := path.Ext(file.Filename)              //文件后缀
	imgFileName := ff_random.KrandAll(32) + FileSuffix //上传文件名称

	FileOpen, err := file.Open()
	defer FileOpen.Close()

	imgFileContent, err := ioutil.ReadAll(FileOpen)
	err = ff_upload.ResumeUpload(ff_upload.QiNiuPublicBucket, imgFileName, imgFileContent)
	if err != nil {
		return "", err
	}
	return imgFileName, nil
}

// @Summary 上传文件
// @tags Upload
// @Accept  multipart/form-data
// @Produce  json
// @Param Token header string true "Token"
// @Param file formData file true "文件"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /upload/file [post]
func UploadFile(ctx *gin.Context) {
	var code int
	file, err := ctx.FormFile("file")
	if err != nil {
		logrus.Errorf("Upload file. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "文件上传失败", "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	date := time.Now().Format("2006-01-02")
	randStr := ff_random.KrandLowerChar(10)
	dir := fmt.Sprintf("./upload/merchant-%v/%v", merchantId, date)
	exist, _ := PathExists(dir)
	if !exist {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			logrus.Errorf("Create dir failure. Error : %v", err)
			code = ff_err.ERROR
			ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "文件上传失败", "data": ""})
			return
		}
	}

	filePath := fmt.Sprintf("%v/%v-%v", dir, randStr, file.Filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		logrus.Errorf("Save file. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "文件上传失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "文件上传成功", "data": map[string]string{"file_url": filePath}})
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
