package ff_upload

import (
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"time"
)

//上传目录
const UPLOAD_FILE_DIR = "E:\\test\\"

/**
上传示例：
formFile,err := ctx.FormFile("img")
if err != nil {
	println(err.Error())
}
filePath,_ := ff_upload.UploadFile(formFile)
println(filePath)
*/

//上传文件
func UploadFile(formFile *multipart.FileHeader) (string, error) {
	file, err := formFile.Open()
	defer file.Close()
	if err != nil {
		return "", err
	}
	//检测目录是否存在
	if !CheckFileIsExist(UPLOAD_FILE_DIR) {
		err := os.Mkdir(UPLOAD_FILE_DIR, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	fileSuffix := path.Ext(formFile.Filename)                      //文件后缀
	filePath := UPLOAD_FILE_DIR + GetRandomString(32) + fileSuffix //文件路径
	fh, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer fh.Close()
	io.Copy(fh, file)

	return filePath, nil
}

//检测目录是否存在
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		println(filename + " not exist")
		exist = false
	}
	return exist
}

//生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
