package ff_header

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"dollmachine/dollrpc/ff_config/ff_const"
	"dollmachine/dollrpc/ff_core/ff_common/ff_random"
	"dollmachine/dollrpc/ff_core/ff_repository/ff_redis/ff_session"
	"strconv"
	"strings"
	"time"
)

type CCHeader struct {
	Context *fasthttp.RequestCtx
}

func NewCCHeader(ctx *fasthttp.RequestCtx) *CCHeader {
	return &CCHeader{Context: ctx}
}

func (p *CCHeader) GetRequestServiceSession() string {
	if p.Context.Request.Header.Peek(ff_const.CCHeaderServiceSession) != nil {
		return string(p.Context.Request.Header.Peek(ff_const.CCHeaderServiceSession))
	}
	return ""
}

func (p *CCHeader) GetRequestReqId() string {
	if p.Context.Request.Header.Peek(ff_const.CCHeaderReqId) != nil {
		return string(p.Context.Request.Header.Peek(ff_const.CCHeaderReqId))
	}
	return ""
}

func (p *CCHeader) SetRequestReqId() {
	if p.Context.Request.Header.Peek(ff_const.CCHeaderReqId) == nil {
		var buf bytes.Buffer
		buf.WriteString("ReqId-")
		buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
		buf.WriteString("-")
		buf.WriteString(ff_random.KrandAll(16))
		p.Context.Request.Header.Set(ff_const.CCHeaderReqId, buf.String())
		p.Context.SetUserValue(ff_const.CCHeaderReqId, buf.String())
	}
}

func (p *CCHeader) SetRequestUserKey(userKey *ff_session.UserKey) {
	p.Context.SetUserValue(ff_const.CCHeaderUserKey, userKey)
}

func (p *CCHeader) GetRequestUserKey() *ff_session.UserKey {
	userKey := p.Context.UserValue(ff_const.CCHeaderUserKey)
	if userKey == nil {
		return nil
	}
	return p.Context.UserValue(ff_const.CCHeaderUserKey).(*ff_session.UserKey)
}

func (p *CCHeader) GetRequestUserKeyMerchantId() int64 {
	userKey := p.GetRequestUserKey()
	if userKey == nil {
		return 0
	}
	return userKey.MerchantId
}

func (p *CCHeader) GetRequestUserKeyUserId() int64 {
	userKey := p.GetRequestUserKey()
	if userKey == nil {
		return 0
	}
	return userKey.UserId
}

func (p *CCHeader) SetResponseReqId() {
	reqId := p.Context.Request.Header.Peek(ff_const.CCHeaderReqId)
	if reqId != nil {
		p.Context.Response.Header.SetBytesV(ff_const.CCHeaderReqId, reqId)
	}
}

func (p *CCHeader) getRequestProxyIp() []string {
	if ips := p.Context.Request.Header.Peek("X-Forwarded-For"); ips != nil {
		return strings.Split(string(ips), ",")
	}
	return []string{}
}

func (p *CCHeader) GetRequestIp() string {
	ips := p.getRequestProxyIp()
	if len(ips) > 0 && ips[0] != "" {
		rip := strings.Split(ips[0], ":")
		return rip[0]
	}
	ip := strings.Split(p.Context.RemoteAddr().String(), ":")
	if len(ip) > 0 {
		if ip[0] != "[" {
			return ip[0]
		}
	}
	return "127.0.0.1"
}

func (p *CCHeader) SetRequestClientIp(clientIp string) {
	p.Context.Request.Header.Add(ff_const.CCHeaderClientIp, clientIp)
}

func (p *CCHeader) SetRequestPlfUserKey(userKey *ff_session.PlfUserKey) {
	p.Context.SetUserValue(ff_const.CCHeaderPlfUserKey, userKey)
}

func (p *CCHeader) GetRequestPlfUserKey() *ff_session.PlfUserKey {
	userKey := p.Context.UserValue(ff_const.CCHeaderPlfUserKey)
	if userKey == nil {
		return nil
	}
	return p.Context.UserValue(ff_const.CCHeaderPlfUserKey).(*ff_session.PlfUserKey)
}

func (p *CCHeader) SetRequestMchUserKey(userKey *ff_session.MchUserKey) {
	p.Context.SetUserValue(ff_const.CCHeaderMchUserKey, userKey)
}

func (p *CCHeader) GetRequestMchUserKey() *ff_session.MchUserKey {
	userKey := p.Context.UserValue(ff_const.CCHeaderMchUserKey)
	if userKey == nil {
		return nil
	}
	return p.Context.UserValue(ff_const.CCHeaderMchUserKey).(*ff_session.MchUserKey)
}
