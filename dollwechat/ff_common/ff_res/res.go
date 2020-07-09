package ff_res

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"dollmachine/dollwechat/ff_common/ff_json"
	"dollmachine/dollwechat/ff_common/ff_page"
	"dollmachine/dollwechat/ff_config/ff_const"
	"net/http"
	"strings"
)

type CCRes struct {
	Code    int64         `json:"code"`
	SubCode string        `json:"sub_code"`
	Msg     string        `json:"msg"`
	ReqId   string        `json:"req_id,omitempty"`
	Data    interface{}   `json:"data,omitempty"`
	Page    *ff_page.Page `json:"page,omitempty"`
}

type CCErr struct {
	Code      int64  `json:"code"`
	SubCode   string `json:"sub_code"`
	Msg       string `json:"msg"`
	SystemErr string `json:"system_err"`
}

func NewCCErr(showMsg string, code int64, code1Type string, code2Type string, code3TypeExtra string, code3Type string, systemErr ...string) *CCErr {
	ccErr := &CCErr{}
	var buf bytes.Buffer
	buf.WriteString("isv.")
	buf.WriteString(code1Type)
	buf.WriteString(code2Type)
	if code3TypeExtra != "" {
		buf.WriteString(code3TypeExtra)
		buf.WriteString(":")
	}
	buf.WriteString(code3Type)
	ccErr.Code = code
	ccErr.SubCode = buf.String()
	ccErr.Msg = showMsg
	ccErr.SystemErr = strings.Join(systemErr, "")
	return ccErr
}

func newCCFail(reqId string, ccErr *CCErr) *CCRes {
	ccRes := &CCRes{
		ReqId: reqId,
	}
	ccRes.Code = ccErr.Code
	ccRes.SubCode = ccErr.SubCode
	ccRes.Msg = ccErr.Msg
	if ccErr.SystemErr != "" {
		log.WithField("err", "newCCFail").WithField("reqId", reqId).Error(ccErr.SystemErr)
	}
	ccRes.Page = nil
	return ccRes
}

func newCCSuccess(reqId string, data interface{}, page *ff_page.Page) *CCRes {
	ccRes := &CCRes{
		Code:    0,
		SubCode: "",
		Msg:     "success",
		ReqId:   reqId,
		Data:    data,
		Page:    page,
	}
	return ccRes
}

func NetHttpNewFailJsonResp(r *http.Request, w http.ResponseWriter, ccErr *CCErr) {
	respFail := newCCFail(r.Header.Get(ff_const.CCHeaderReqId), ccErr)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write([]byte(ff_json.MarshalToStringNoError(respFail)))
}

func NetHttpNewSuccessNormalJsonResp(r *http.Request, w http.ResponseWriter) {
	NetHttpNewSuccessPageJsonResp(r, w, map[string]string{"result": "success"}, nil)
}

func NetHttpNewSuccessJsonResp(r *http.Request, w http.ResponseWriter, data interface{}) {
	NetHttpNewSuccessPageJsonResp(r, w, data, nil)
}

func NetHttpNewSuccessPageJsonResp(r *http.Request, w http.ResponseWriter, data interface{}, page *ff_page.Page) {
	respSuccess := newCCSuccess(r.Header.Get(ff_const.CCHeaderReqId), data, page)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write([]byte(ff_json.MarshalToStringNoError(respSuccess)))
}
