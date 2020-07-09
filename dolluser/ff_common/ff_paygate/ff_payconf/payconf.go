package ff_payconf

import (
	"errors"
)

const (
	WxPreCreate  = "https://spay.fuiou.com/wxPreCreate" //微信公众号、微信小程序、微信APP，支付宝服务窗等
	PreCreate    = "https://spay.fuiou.com/preCreate"   //主扫统一下单
	CommonQuery  = "https://spay.fuiou.com/commonQuery" //订单查询
	SubAppid     = "wx33c1d48cde2330e4"
	FyAppid      = "wxfa089da95020ba1a"
	NotifyEspUrl = "https://wawafront.tunnel.aioil.cn/ks/v1/oauth/recharge/notify_esp"
	NotifyUrl    = "https://wawafront.tunnel.aioil.cn/ks/v1/oauth/recharge/notify"

	WechatPay_Mini_FuYou   = "wechatpay_mini_fy"
	WechatPay_Qr_FuYou     = "wechatpay_qr_fy"
	WECHATPAY_LETPAY_FUYOU = "wechatpay_letpay_fy"
	WECHATPAY_JSPAY_FUYOU  = "wechatpay_jspay_fy"
	FUYOU_MCH_ID           = "0005840F1382995" //"0002900F0370542" // "0002900F0370588"
	FUYOU_MCH_PLATFORM     = "08M0025639"
)

var CHANNEL_ALLOW_MAP = map[string]string{
	WECHATPAY_LETPAY_FUYOU: "1",
	WechatPay_Qr_FuYou:     "2",
	WechatPay_Mini_FuYou:   "3",
	WECHATPAY_JSPAY_FUYOU:  "4",
}

type PayConf struct {
	DefaultMeta string
	Meta        map[string]*PayConfMeta
}

type PayConfMeta struct {
	Channel     string
	MchPlatform string //第三方机构号
	MchId       string //第三方的商户号
	Cert        string //证书
	Secret      string //加密串
	RsaPrivate  []byte //密钥
	RsaPublic   []byte //公钥
}

func NewPayConf() *PayConf {
	payConf := &PayConf{}
	payConf.DefaultMeta = FUYOU_MCH_ID
	payConf.Meta = make(map[string]*PayConfMeta)
	payConf.Meta[FUYOU_MCH_ID] = &PayConfMeta{
		Channel:     WECHATPAY_JSPAY_FUYOU,
		MchPlatform: FUYOU_MCH_PLATFORM,
		MchId:       FUYOU_MCH_ID,
		Cert:        "",
		Secret:      "",
		RsaPrivate: []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAIggsGdOn9uvmN5sFFGw5jPrSqE6
KMOJtUZUsZ51oIMn8sjn7XyradtiArFCjaTPZkrPSxuU3+f2XoUB/dfTbQGOzqM+qDnWvdPc4UJj
6FsE/XQuARETIM8dDnbJZBcfgD7Xpc+hLfKfC/FfuI6HQZpwQGao25VmhtS9Ji3JJGiHAgMBAAEC
gYBI+8W4TZY1kYrTTX0DY2W41mDX2je6xp0zDPuB6qzZRNTNVFOmsLx7i6vH39fTUgMU/tjU+9ek
JRn+E9hGG6voCmcYHfzkfxQMMdajWl6nkPq1kbvVoA8OJe1JaW7lNq/OuriEOU876IAwcWbOLeMq
rGuB3v/JrNq5EY3KJyrzaQJBAOXKBNOU8jDKjJnSdzALhluRZe1ftY1WS2aYMui4lfQO3hUKPbRA
K9PjUVib049IVG1xlqhWtSMdK1BK6W3EJ5sCQQCXp7Mb92TKoVN5Rbv9ZsXJdKJiKetYLV3EwsGF
0G5w571Qka8C9SfKevvcTnksUcbidwvJDDYmiytsU2ABuU+FAkAw7gJ3Fzk3AHpN6s3sUhfq+Zvt
nrqm/OATWYdFnMB5doz9h++5qQxsEvRoXM4ArZMkttIwyD3L21M0xq7L67/PAkAj3hnSZ3SDKByh
9gg8Km5k8xzksp1iwXgH7Tfv+hfkxCpWP95wiKLclLG0rSqjfMPZE+bJqgW0n/2pJR7zyWwxAkAG
j/0zZSmDL1AsQwGvN3xDDFw0XwNDWQ1b+ZTMyCH5CWGKbxLQwZZJSN33q2LIck8VwJ9FUqtVlv8+
vQ49XjkS
-----END RSA PRIVATE KEY-----`),
		RsaPublic: []byte(`MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDz2fCOYaaU6sztFql4cOmiFRq2LRk1XuGfrJnMFa09QMXMXOEn9YNYC44zV1AE/q9b0BKGbM74YPoge/7qsW+Heao76Drv6HujP+rXLFbsXT5f9rcID2GCzDc+DXjb+NfwSa8vS9KJ3dau2xm87zpjdQ9zER6VH4UcZTgj7LbzgwIDAQAB`),
	}
	return payConf
}

func (p *PayConf) GetDefaultPayConfMeta() (*PayConfMeta, error) {
	return p.GetPayConfMeta(p.DefaultMeta)
}

func (p *PayConf) GetPayConfMeta(meta string) (*PayConfMeta, error) {
	payConfMeta, ok := p.Meta[meta]
	if !ok {
		return nil, errors.New("支付渠道不存在")
	}
	return payConfMeta, nil
}

func (p *PayConfMeta) SetChannel(channel string) error {
	if _, ok := CHANNEL_ALLOW_MAP[channel]; !ok {
		return errors.New("支付渠道类型不存在")
	}

	p.Channel = channel

	return nil
}
