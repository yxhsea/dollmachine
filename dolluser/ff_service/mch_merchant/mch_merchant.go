package mch_merchant

import (
	"dollmachine/dolluser/ff_model/mch_merchant"
)

type MerchantService struct {
}

func NewMerchantService() *MerchantService {
	return &MerchantService{}
}

func (p *MerchantService) CheckIsExitsByMchMerchantId(merchantId int64) bool {
	mchMerchant := mch_merchant.NewMchMerchant()
	return mchMerchant.CheckIsExitsByMerchantId(merchantId)
}
