package ff_c50x

import (
	log "github.com/sirupsen/logrus"
	"dollmachine/dollwechat/ff_common/ff_res"
	"net/http"
)

func ApiServerPanic(w http.ResponseWriter, r *http.Request, ps interface{}){
	fields := make(log.Fields)
	fields["c50x"] = "ApiServerPanic"
	fields["client-ip"] = r.RemoteAddr
	fields["request-url"] = r.URL
	log.WithFields(fields).Error("ApiServerPanic")
	ccErr := ff_res.NewCCErr("接口发生未知错误", ff_res.ErrorCode0ServiceMethodNotFoundError, ff_res.ErrorCode1TypeInvalidRequestErr, ff_res.ErrorCode2TypeUnknownError, "api", ff_res.ErrorCode3TypeUnexpectedError, "ApiServerPanic")
	ff_res.NetHttpNewFailJsonResp(r, w, ccErr)
	return
}
