package ff_c50x

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"dollmachine/dollwechat/ff_common/ff_res"
)

func NotFound(rw http.ResponseWriter, r *http.Request) {
	fields := make(log.Fields)
	fields["c50x"] = "NotFound"
	fields["client-ip"] = r.RemoteAddr
	fields["request-url"] = r.URL
	log.WithFields(fields).Error("NotFound")
	ccErr := ff_res.NewCCErr("接口不存在", ff_res.ErrorCode0ServiceMethodNotFoundError, ff_res.ErrorCode1TypePlatformApiErr, ff_res.ErrorCode2TypeInvalidParameter, "api", ff_res.ErrorCode3TypeNotExist, "NotFound")
	ff_res.NetHttpNewFailJsonResp(r, rw, ccErr)
	return
}
