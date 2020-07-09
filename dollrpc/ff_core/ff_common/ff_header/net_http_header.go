package ff_header

import (
	"bytes"
	"context"
	"dollmachine/dollrpc/ff_config/ff_const"
	"dollmachine/dollrpc/ff_core/ff_common/ff_random"
	"dollmachine/dollrpc/ff_core/ff_repository/ff_redis/ff_session"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type NetHttpHeader struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

func NewNetHttpHeader(request *http.Request, responseWriter http.ResponseWriter) *NetHttpHeader {
	return &NetHttpHeader{
		Request:        request,
		ResponseWriter: responseWriter,
	}
}

func (p *NetHttpHeader) GetRequestServiceSession() string {
	return p.Request.Header.Get(ff_const.CCHeaderServiceSession)
}

func (p *NetHttpHeader) GetRequestReqId() string {
	return p.Request.Header.Get(ff_const.CCHeaderReqId)
}

func (p *NetHttpHeader) SetRequestReqId() {
	if p.Request.Header.Get(ff_const.CCHeaderReqId) == "" {
		var buf bytes.Buffer
		buf.WriteString("ReqId-")
		buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
		buf.WriteString("-")
		buf.WriteString(ff_random.KrandAll(16))
		p.Request.Header.Set(ff_const.CCHeaderReqId, buf.String())

		c := p.Request.Context()
		context.WithValue(c, ff_const.CCHeaderReqId, buf.String())
		p.Request.WithContext(c)
	}
}

func (p *NetHttpHeader) SetRequestUserKey(userKey *ff_session.UserKey) {
	c := p.Request.Context()
	context.WithValue(c, ff_const.CCHeaderUserKey, userKey)
	p.Request.WithContext(c)
}

func (p *NetHttpHeader) GetRequestUserKey() *ff_session.UserKey {
	userKey := p.Request.Context().Value(ff_const.CCHeaderUserKey)
	if userKey == nil {
		return nil
	}
	return p.Request.Context().Value(ff_const.CCHeaderUserKey).(*ff_session.UserKey)
}

func (p *NetHttpHeader) GetRequestUserKeyMerchantId() int64 {
	userKey := p.GetRequestUserKey()
	if userKey == nil {
		return 0
	}
	return userKey.MerchantId
}

func (p *NetHttpHeader) GetRequestUserKeyUserId() int64 {
	userKey := p.GetRequestUserKey()
	if userKey == nil {
		return 0
	}
	return userKey.UserId
}

func (p *NetHttpHeader) SetResponseReqId() {
	reqId := p.Request.Header.Get(ff_const.CCHeaderReqId)
	if reqId != "" {
		p.ResponseWriter.Header().Set(ff_const.CCHeaderReqId, reqId)
	}
}

func (p *NetHttpHeader) getRequestProxyIp() []string {
	if ips := p.Request.Header.Get("X-Forwarded-For"); ips != "" {
		return strings.Split(ips, ",")
	}
	return []string{}
}

func (p *NetHttpHeader) GetRequestIp() string {
	ips := p.getRequestProxyIp()
	if len(ips) > 0 && ips[0] != "" {
		rip := strings.Split(ips[0], ":")
		return rip[0]
	}
	ip := strings.Split(p.Request.RemoteAddr, ":")
	if len(ip) > 0 {
		if ip[0] != "[" {
			return ip[0]
		}
	}
	return "127.0.0.1"
}

func (p *NetHttpHeader) SetRequestClientIp(clientIp string) {
	p.Request.Header.Set(ff_const.CCHeaderClientIp, clientIp)
}

func (p *NetHttpHeader) SetRequestPlfUserKey(userKey *ff_session.PlfUserKey) {
	c := p.Request.Context()
	context.WithValue(c, ff_const.CCHeaderPlfUserKey, userKey)
	p.Request.WithContext(c)
}

func (p *NetHttpHeader) GetRequestPlfUserKey() *ff_session.PlfUserKey {
	userKey := p.Request.Context().Value(ff_const.CCHeaderPlfUserKey)
	if userKey == nil {
		return nil
	}
	return p.Request.Context().Value(ff_const.CCHeaderPlfUserKey).(*ff_session.PlfUserKey)
}

func (p *NetHttpHeader) SetRequestMchUserKey(userKey *ff_session.MchUserKey) {
	c := p.Request.Context()
	context.WithValue(c, ff_const.CCHeaderMchUserKey, userKey)
	p.Request.WithContext(c)
}

func (p *NetHttpHeader) GetRequestMchUserKey() *ff_session.MchUserKey {
	userKey := p.Request.Context().Value(ff_const.CCHeaderMchUserKey)
	if userKey == nil {
		return nil
	}
	return p.Request.Context().Value(ff_const.CCHeaderMchUserKey).(*ff_session.MchUserKey)
}
