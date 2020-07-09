package ff_header

import (
	"github.com/gin-gonic/gin"
	"dollmachine/dolluser/ff_cache/user_session"
)

type FFHeader struct {
	Context *gin.Context
}

func NewFFHeader(ctx *gin.Context) *FFHeader {
	return &FFHeader{Context: ctx}
}

func (p *FFHeader) GetRequestUserSession() *user_session.UserSession {
	token := p.Context.Request.Header.Get("Token")
	return user_session.NewUserSession().GetUserSession(token)
}
