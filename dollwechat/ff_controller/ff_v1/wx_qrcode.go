package ff_v1

import (
	"dollmachine/dollwechat/ff_config/ff_const"
	"dollmachine/dollwechat/ff_common/ff_random"
	"gopkg.in/chanxuehong/wechat.v2/mp/qrcode"
	"dollmachine/dollwechat/ff_config/ff_vars"
	"dollmachine/dollwechat/ff_common/ff_res"
	log "github.com/sirupsen/logrus"
	"github.com/gogap/logrus"
	"github.com/gin-gonic/gin"
)

func GetWxQrCode(ctx *gin.Context) {
	sceneStr := ff_const.FFUserLoginQrCodeKey + "-" + ff_random.KrandAll(6)
	tempQrCode, err := qrcode.CreateStrSceneTempQrcode(ff_vars.WxMpWechatClient, sceneStr, 3600)
	logrus.Debugf("Create TempQrCode. sceneStr : %v", sceneStr)
	if err != nil {
		log.Error("Create TempQrCode Fail Error : " + err.Error())
		ccErr := ff_res.NewCCErr(ff_res.ErrorMsgThirdServiceError,ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError,"创建二维码失败," + err.Error())
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}

	ff_res.NetHttpNewSuccessJsonResp(ctx.Request, ctx.Writer, map[string]interface{}{"qr_code":tempQrCode.URL,"scene_str":sceneStr})
	return
}