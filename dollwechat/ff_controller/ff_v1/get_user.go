package ff_v1

import (
	"dollmachine/dollwechat/ff_config/ff_vars"
	"dollmachine/dollwechat/ff_redis"
	log "github.com/sirupsen/logrus"
	"dollmachine/dollwechat/ff_common/ff_res"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//查询用户
func GetLoginUser(ctx *gin.Context){
	sceneStr := ctx.PostForm("scene_str")

	if len(sceneStr) <= 0 {
		ccErr := ff_res.NewCCErr(ff_res.ErrorMsgParameterError,ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError,"sceneStr参数不能为空")
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}

	conn := ff_vars.RedisConn.Get()
	defer conn.Close()
	flag, err := ff_redis.NewString().Exists(conn,sceneStr)
	if err != nil {
		log.Errorf("query scenestr fail Error : %s ", err.Error())
		ccErr := ff_res.NewCCErr(ff_res.ErrorMsgCacheError,ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError,"查询信息失败")
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}
	if !flag {
		ccErr := ff_res.NewCCErr(ff_res.ErrorMsgCacheError,ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError,"所查用户不存在")
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}
	userInfo, err := ff_redis.NewString().Get(conn, sceneStr)
	if err != nil {
		ccErr := ff_res.NewCCErr(ff_res.ErrorMsgCacheError,ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError, err.Error())
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}

	var res map[string]interface{}
	err = json.Unmarshal([]byte(userInfo), &res)
	if err != nil {
		log.Warnf("userInfo to map fail : %s ", err.Error())
	}
	ff_res.NetHttpNewSuccessJsonResp(ctx.Request, ctx.Writer, res)
	return
}