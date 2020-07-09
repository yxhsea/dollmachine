package ff_header

import (
	"github.com/gin-gonic/gin"
	"dollmachine/dollplatform/ff_cache/platform_session"
)

type FFHeader struct {
	Context *gin.Context
}

func NewFFHeader(ctx *gin.Context) *FFHeader {
	return &FFHeader{Context: ctx}
}

func (p *FFHeader) GetRequestPlatformSession() *platform_session.PlatformSession {
	accessToken := p.Context.Request.Header.Get("Token")
	return platform_session.NewPlatformSession().GetPlatformSession(accessToken)
}
