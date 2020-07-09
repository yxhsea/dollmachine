package ff_header

import (
	"github.com/gin-gonic/gin"
	"dollmachine/dollmerchant/ff_cache/merchant_session"
)

type FFHeader struct {
	Context *gin.Context
}

func NewFFHeader(ctx *gin.Context) *FFHeader {
	return &FFHeader{Context: ctx}
}

func (p *FFHeader) GetRequestMerchantSession() *merchant_session.MerchantSession {
	accessToken := p.Context.Request.Header.Get("Token")
	return merchant_session.NewMerchantSession().GetMerchantSession(accessToken)
}
