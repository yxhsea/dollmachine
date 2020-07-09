package ff_v1

import (
	"fmt"
	"dollmachine/dollwechat/ff_config/ff_vars"
	"gopkg.in/chanxuehong/wechat.v2/mp/menu"
	"github.com/gin-gonic/gin"
)

func CreateMenu(ctx *gin.Context) {
	MenuList := &menu.Menu{
		Buttons: []menu.Button{
			menu.Button{
				Type:       "view",                                                     // 非必须; 菜单的响应动作类型
				Name:       "设备绑定",                                                     // 必须;  菜单标题
				Key:        "",                                                         // 非必须; 菜单KEY值, 用于消息接口推送
				URL:        "http://wawadevicebind.tunnel.aioil.cn/dist/#/device_list", // 非必须; 网页链接, 用户点击菜单可打开链接
				MediaId:    "",                                                         // 非必须; 调用新增永久素材接口返回的合法media_id
				AppId:      "",                                                         // 非必须; 跳转到小程序的appid
				PagePath:   "",                                                         // 非必须; 跳转到小程序的path
				SubButtons: []menu.Button{},                                            // 非必须; 二级菜单数组
			},
			menu.Button{
				Type:       "view",                                                                                                  // 非必须; 菜单的响应动作类型
				Name:       "抓娃娃",                                                                                                   // 必须;  菜单标题
				Key:        "",                                                                                                      // 非必须; 菜单KEY值, 用于消息接口推送
				URL:        "https://wawafront.tunnel.aioil.cn/dist/?from=singlemessage&isappinstalled=0#/?merchant_id=20000332793", // 非必须; 网页链接, 用户点击菜单可打开链接
				MediaId:    "",                                                                                                      // 非必须; 调用新增永久素材接口返回的合法media_id
				AppId:      "",                                                                                                      // 非必须; 跳转到小程序的appid
				PagePath:   "",                                                                                                      // 非必须; 跳转到小程序的path
				SubButtons: []menu.Button{},                                                                                         // 非必须; 二级菜单数组
			},
			menu.Button{
				Type:       "view",                                                                                                                     // 非必须; 菜单的响应动作类型
				Name:       "测试房间",                                                                                                                     // 必须;  菜单标题
				Key:        "",                                                                                                                         // 非必须; 菜单KEY值, 用于消息接口推送
				URL:        "https://wawafront.tunnel.aioil.cn/dist/?from=singlemessage#/game?room_id=26020266&device_id=9999&merchant_id=20000332793", // 非必须; 网页链接, 用户点击菜单可打开链接
				MediaId:    "",                                                                                                                         // 非必须; 调用新增永久素材接口返回的合法media_id
				AppId:      "",                                                                                                                         // 非必须; 跳转到小程序的appid
				PagePath:   "",                                                                                                                         // 非必须; 跳转到小程序的path
				SubButtons: []menu.Button{},                                                                                                            // 非必须; 二级菜单数组
			},
		},
	}
	err := menu.Create(ff_vars.WxMpWechatClient, MenuList)
	if err != nil {
		fmt.Println("______", err.Error())
	}
}
