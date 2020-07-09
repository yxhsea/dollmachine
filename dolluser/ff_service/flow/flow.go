package flow

import (
	"errors"
	"dollmachine/dolluser/ff_common/ff_page"
	"dollmachine/dolluser/ff_model/flow_record"
)

type FlowService struct {
}

func NewFlowService() *FlowService {
	return &FlowService{}
}

func (p *FlowService) GetFlowList(merchantId int64, offset int, pageSize int, totalSize int) ([]map[string]interface{}, *ff_page.Page, error) {
	page := ff_page.NewPage(offset, pageSize, totalSize)
	flowDao := flow_record.NewFlowRecord()
	FlowList, err := flowDao.GetFlowRecordListByMerchantId(merchantId, offset, pageSize, "*")
	if err != nil || len(FlowList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := flowDao.GetFlowRecordListTotalCount(merchantId)
	page.SetTotalSize(count)
	return FlowList, page, nil
}
