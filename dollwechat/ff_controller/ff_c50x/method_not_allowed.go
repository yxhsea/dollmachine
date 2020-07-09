package ff_c50x

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"dollmachine/dollwechat/ff_common/ff_res"
)

func MethodNotAllowed(w http.ResponseWriter, r *http.Request){
	fields := make(log.Fields)
	fields["c50x"] = "MethodNotAllowed"
	fields["client-ip"] = r.RemoteAddr
	fields["request-url"] = r.URL
	log.WithFields(fields).Error("MethodNotAllowed")
	ccErr := ff_res.NewCCErr("接口不允许调用", ff_res.ErrorCode0ServiceMethodNotAllowedError, ff_res.ErrorCode1TypeInvalidRequestErr, ff_res.ErrorCode2TypeInvalidParameter, "api", ff_res.ErrorCode3TypeCallLimited, "MethodNotAllowed")
	ff_res.NetHttpNewFailJsonResp(r, w, ccErr)
	return
}
