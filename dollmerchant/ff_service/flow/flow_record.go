package flow

import (
	"errors"
	"fmt"
	"github.com/Unknwon/com"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_model/flow_record"
	"dollmachine/dollmerchant/ff_model/usr_user"
)

type FlowRecordService struct {
}

func NewFlowRecordService() *FlowRecordService {
	return &FlowRecordService{}
}

func (p *FlowRecordService) GetFlowRecordList(offset, pageSize, totalSize, userId, iType string, merchantId, StartTime, EndTime int64) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	UserId, _ := com.StrTo(userId).Int64()
	IType, _ := com.StrTo(iType).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	flowRecordDao := flow_record.NewFlowRecord()
	flowRecordList, err := flowRecordDao.GetFlowRecordList(Offset, PageSize, "*", merchantId, StartTime, EndTime, UserId, IType)
	if err != nil || len(flowRecordList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	//补齐用户名
	usrDao := usr_user.NewUsrUser()
	for _, v := range flowRecordList {
		usrInfo := usrDao.GetUsrUserOne(fmt.Sprint(v["user_id"]), "nick_name")
		v["user_name"] = usrInfo["nick_name"]
	}

	count, _ := flowRecordDao.GetFlowRecordListTotalCount(merchantId, StartTime, EndTime, UserId, IType)
	page.SetTotalSize(count)
	return flowRecordList, page, nil
}
