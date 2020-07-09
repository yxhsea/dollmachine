package ff_upload

import (
	"testing"
	//"fmt"cls
	//"dollmachine/dollrpc/ff_config/ff_vars"
	//"io/ioutil"
	/*"fmt"
	"dollmachine/dollrpc/ff_config/ff_vars"*/
	//	"io/ioutil"
	"fmt"
	"github.com/qiniu/api.v7/storage"
	"dollmachine/dollrpc/ff_config/ff_vars"
)

/*func TestDownloadFile(t *testing.T) {
	err := DownloadFile(GetPrivateFileLink("http://p1xjgl9re.bkt.clouddn.com", "t/a.mpg", 0), "kael_test.mpg")
	if err != nil {
		fmt.Println("TestDownloadFile", err.Error())
	}else {
		fmt.Println("TestDownloadFile success",)
	}

}*/

/*
func TestGetPrivateFileLink(t *testing.T) {
*/
/*access_key = "AYMygeviPY0F5dNp2yhJJ4cEQUXcVA4ozhnwnMMJ"
secret_key = "-oB53OIF1DU_il2gVx_HAHquyzm05iZgiw8Zyz4D"*/ /*

	ff_vars.QiNiuAccessKey = "AYMygeviPY0F5dNp2yhJJ4cEQUXcVA4ozhnwnMMJ"
	ff_vars.QiNiuSecretKey = "-oB53OIF1DU_il2gVx_HAHquyzm05iZgiw8Zyz4D"
	//fmt.Println("TestGetPrivateFileLink", GetPrivateFileLink("http://p1xjgl9re.bkt.clouddn.com", "t/a.mpg", 0))
	fmt.Println("TestGetPrivateFileLink", GetPrivateFileLink("http://p1xjgl9re.bkt.clouddn.com", "wx_logo.jpg", 0))
	//tmpVideo := "tmp.mp4"
	//DownloadFile(GetPrivateFileLink("http://p1xjgl9re.bkt.clouddn.com", "0OuMIhgy3b36IFIO5a71YhmWSg36W0GP.mp4", 0),tmpVideo)
}
*/

func TestUpload(t *testing.T) {
	ff_vars.QiNiuAccessKey = "AYMygeviPY0F5dNp2yhJJ4cEQUXcVA4ozhnwnMMJ"
	ff_vars.QiNiuSecretKey = "-oB53OIF1DU_il2gVx_HAHquyzm05iZgiw8Zyz4D"
	ff_vars.QiNiuZone = &storage.ZoneHuanan

	/*inputFile := "E:\\tupian.jpg"
	fileContent, err := ioutil.ReadFile(inputFile)
	if err != nil {
		println(err.Error())
	}
	bucket := "kael-test"
	err = ResumeUpload(bucket,"2.jpg",fileContent)
	if err != nil {
		println(err.Error())
	}
	println(6666)*/

	fmt.Println("TestGetPublicFileLink", GetPublicFileLink("http://p21gswpth.bkt.clouddn.com", "I7e9166tu7wCZ9b7G7w7El3ChSIeNWau.jpg"))
	fmt.Println("TestGetPrivateFileLink", GetPrivateFileLink("http://p1xjgl9re.bkt.clouddn.com", "I7e9166tu7wCZ9b7G7w7El3ChSIeNWau.jpg", 0))
}
