package ff_upload

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"dollmachine/dollmerchant/ff_common/ff_json"
	"dollmachine/dollmerchant/ff_common/ff_random"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const QiNiuPrivateBucket = "kael-test"
const QiNiuPublicBucket = "kael-test-public"
const QiNiuPrivateDomain = "http://p1xjgl9re.bkt.clouddn.com"
const QiNiuPublicDomain = "http://p21gswpth.bkt.clouddn.com"

// 指定的进度文件保存目录，实际情况下，请确保该目录存在，而且只用于记录进度文件
const RECORD_PROCESS_DIR = "E:\\download\\"

// 下载文件存放路径
//const LOCAL_DOWNLOAD_RID = "/data/www/qiniu/download/"
const LOCAL_DOWNLOAD_RID = "E:\\download\\"

type ProgressRecord struct {
	Progresses []storage.BlkputRet `json:"progresses"`
}

func md5Hex(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetToken(bucket string) string {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(ff_vars.QiNiuAccessKey, ff_vars.QiNiuSecretKey)
	return putPolicy.UploadToken(mac)
}

func ResumeUpload(bucket string, key string, file []byte) error {
	upToken := GetToken(bucket)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = ff_vars.QiNiuZone
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	fileSize := int64(len(file))

	rndStr := ff_random.KrandAll(32)
	// 必须仔细选择一个能标志上传唯一性的 recordKey 用来记录上传进度
	recordKey := md5Hex(fmt.Sprintf("%s:%s:%s:%s", bucket, key, time.Now().UnixNano(), rndStr)) + ".progress"

	mErr := os.MkdirAll(RECORD_PROCESS_DIR, 0755)
	if mErr != nil {
		fmt.Println("mkdir for record dir error,", mErr)
		return mErr
	}
	recordPath := filepath.Join(RECORD_PROCESS_DIR, recordKey)
	progressRecord := ProgressRecord{}
	// 尝试从旧的进度文件中读取进度
	recordFp, openErr := os.Open(recordPath)
	if openErr == nil {
		progressBytes, readErr := ioutil.ReadAll(recordFp)
		if readErr == nil {
			mErr := ff_json.UnmarshalByte(progressBytes, &progressRecord)
			if mErr == nil {
				// 检查context 是否过期，避免701错误
				for _, item := range progressRecord.Progresses {
					if storage.IsContextExpired(item) {
						fmt.Println(item.ExpiredAt)
						progressRecord.Progresses = make([]storage.BlkputRet, storage.BlockCount(fileSize))
						break
					}
				}
			}
		}
		recordFp.Close()
	}
	if len(progressRecord.Progresses) == 0 {
		progressRecord.Progresses = make([]storage.BlkputRet, storage.BlockCount(fileSize))
	}

	resumeUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	progressLock := sync.RWMutex{}

	putExtra := storage.RputExtra{
		Progresses: progressRecord.Progresses,
		Notify: func(blkIdx int, blkSize int, ret *storage.BlkputRet) {
			progressLock.Lock()
			progressLock.Unlock()
			//将进度序列化，然后写入文件
			progressRecord.Progresses[blkIdx] = *ret
			progressBytes, _ := ff_json.Marshal(progressRecord)
			fmt.Println("write progress file", blkIdx, recordPath)
			wErr := ioutil.WriteFile(recordPath, progressBytes, 0644)
			if wErr != nil {
				fmt.Println("write progress file error,", wErr)
			}
		},
	}
	err := resumeUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(file), fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}

	os.Remove(recordPath)

	fmt.Println(ret.Key, ret.Hash)

	return nil
}

//获取公开空间下载文件链接
func GetPublicFileLink(domain string, key string) string {
	return storage.MakePublicURL(domain, key)
}

//获取私有空间下载文件链接
func GetPrivateFileLink(domain string, key string, deadline int64) string {
	if deadline == 0 {
		deadline = time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	}

	mac := qbox.NewMac(ff_vars.QiNiuAccessKey, ff_vars.QiNiuSecretKey)
	return storage.MakePrivateURL(mac, domain, key, deadline)
}

func DownloadFile(url string, fileName string) error {
	request := gorequest.New()

	resp, _, requestErr := request.Get(url).End()
	if requestErr != nil {
		return errors.New("获取资源失败，网络异常")
	}

	err := os.MkdirAll(LOCAL_DOWNLOAD_RID, 0755)
	if err != nil {
		fmt.Println("mkdir for record dir error,", err)
		return err
	}

	f, err := os.Create(LOCAL_DOWNLOAD_RID + fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	io.Copy(f, resp.Body)

	return nil
}
