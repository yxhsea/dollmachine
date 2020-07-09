package ff_mqttmsg

import "time"

//TODO::服务器订阅的数据包

//基本数据包结构体
type BasePkg struct {
	DeviceId string      `json:"deviceId"` //设备ID, 唯一编码
	Action   string      `json:"action"`   //HEART/STARTUP等操作标志符
	Status   string      `json:"status"`   //改动作是否成功，1表示成功，0表示失败
	Content  interface{} `json:"content"`  //执行后返回的内容,主体内容
}

//设备心跳包结构体
type HeartPkg struct {
	CameraNum     int    `json:"cameraNum"`     //摄像头数量
	CameraWorking int    `json:"cameraWorking"` //正在运行中的摄像头数量
	DeviceID      string `json:"deviceId"`      //设备ID, 唯一编码
	Mac           string `json:"mac"`           //设备网卡的Mac地址
	Msg           string `json:"msg"`           //信息
	Status        int    `json:"status"`        //状态
	Version       string `json:"version"`       //版本号
}

//设备游戏结果数据包结构体
type PlayOverPkg struct {
	EndTime      int64 `json:"endTime"`      //结束时间
	ExpectResult int64 `json:"expectResult"` //期望结果
	IsAward      int   `json:"isAward"`      //真实结果
	UserId       int64 `json:"userId"`       //用户ID
}

//设备上线数据包
type StartUpPkg struct {
}

//设备掉线数据包
type WillDownPkg struct {
}

//TODO::服务器发布的数据包

type CtlPkg struct {
	Imei    string      `json:"deviceId"` //模块IMEI，唯一码
	Action  string      `json:"action"`   //COIN等操作标志符
	En      string      `json:"en"`       //当该值为1时，执行动作；当该值为0时，禁止该动作
	Content interface{} `json:"content"`  //携带的内容（如写rs232，需要在此处填入内容）
	Code    string      `json:"code"`
	ErrType string      `json:"err_type"`
	ErrMsg  string      `json:"err_msg"`
}

const (
	CtlPkgActionError        = "ERROR"
	CtlPkgActionMoveForward  = "MOVE-FORWARD"
	CtlPkgActionMoveLeft     = "MOVE-LEFT"
	CtlPkgActionMoveRight    = "MOVE-RIGHT"
	CtlPkgActionMoveBackward = "MOVE-BACKWARD"
	CtlPkgActionMoveStop     = "MOVE-STOP"
	CtlPkgActionBegin        = "BEGIN"
	CtlPkgActionCatch        = "CATCH"
	CtlPkgActionBind         = "BIND"
)

func NewCtlPkgErr(imei string, action string, code string, errType string, errMsg string) *CtlPkg {
	return &CtlPkg{
		Imei:    imei,
		Action:  action,
		Code:    code,
		ErrType: errType,
		ErrMsg:  errMsg,
	}
}

func NewCtlPkgMove(imei string, direction string, en string) *CtlPkg {
	var action string
	switch direction {
	case "forward":
		action = CtlPkgActionMoveForward
	case "left":
		action = CtlPkgActionMoveLeft
	case "right":
		action = CtlPkgActionMoveRight
	case "backward":
		action = CtlPkgActionMoveBackward
	case "stop":
		action = CtlPkgActionMoveStop
	default:
		action = CtlPkgActionMoveStop
	}

	return &CtlPkg{
		Imei:   imei,
		Action: action,
		En:     en,
		Content: struct {
		}{},
		Code:   MqttCodeNormal,
		ErrMsg: time.Now().String(),
	}
}

type ContentCtlPkgBegin struct {
	UserId    int64 `json:"user_id"`
	GameParam []int `json:"game_param"`
}

func NewContentCtlPkgBegin(userId int64, gameParam []int) *ContentCtlPkgBegin {
	return &ContentCtlPkgBegin{
		UserId:    userId,
		GameParam: gameParam,
	}
}

func NewCtlPkgBegin(imei string, en string, content *ContentCtlPkgBegin) *CtlPkg {
	return &CtlPkg{
		Imei:    imei,
		Action:  CtlPkgActionBegin,
		En:      en,
		Content: content,
		Code:    MqttCodeNormal,
	}
}

func NewCtlPkgCatch(imei string, en string) *CtlPkg {
	return &CtlPkg{
		Imei:   imei,
		Action: CtlPkgActionCatch,
		En:     en,
		Content: struct {
		}{},
		Code: MqttCodeNormal,
	}
}

type ContentCtlPkgBind struct {
	DeviceId interface{} `json:"device_id"`
	//MerchantId			int64		`json:"merchant_id"`
	//MerchantName		string	`json:"merchant_name"`
}

func NewContentCtlPkgBind(deviceId interface{}) *ContentCtlPkgBind {
	return &ContentCtlPkgBind{
		DeviceId: deviceId,
		//MerchantId:merchantId,
		//MerchantName:merchantName,
	}
}

func NewCtlPkgBind(imei string, en string, content *ContentCtlPkgBind) *CtlPkg {
	return &CtlPkg{
		Imei:    imei,
		Action:  CtlPkgActionBind,
		En:      en,
		Content: content,
		Code:    MqttCodeNormal,
	}
}
