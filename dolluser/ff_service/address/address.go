package address

import (
	"github.com/pkg/errors"
	"dollmachine/dolluser/ff_model/usr_address"
	"time"
)

type AddressService struct {
}

func NewAddressService() *AddressService {
	return &AddressService{}
}

func (p *AddressService) GetAddressInfo(addressId int64) (map[string]interface{}, error) {
	//地址信息
	AddressDao := usr_address.NewUsrAddress()
	if !AddressDao.CheckIsExitsByAddressId(addressId) {
		return nil, errors.New("地址不存在")
	}
	AddressInfo, err := AddressDao.GetAddressInfoByAddressId(addressId, "*")
	if err != nil {
		return nil, err
	}
	return AddressInfo, nil
}

func (p *AddressService) AddAddress(addressId, userId, phone int64, fullName, province, city, district, address string) (int64, error) {
	addressDao := usr_address.NewUsrAddress()
	nowTime := time.Now().Unix()
	var addrId int64
	if addressId > 0 {
		err := addressDao.UpdateAddressInfo(map[string]interface{}{
			"user_id":   userId,
			"phone":     phone,
			"full_name": fullName,
			"province":  province,
			"city":      city,
			"district":  district,
			"address":   address,
			"update_at": nowTime,
		}, addressId)
		if err != nil {
			return 0, err
		}
		addrId = addressId
	} else {
		addressId, err := addressDao.AddAddressInfo(map[string]interface{}{
			"user_id":    userId,
			"phone":      phone,
			"full_name":  fullName,
			"province":   province,
			"city":       city,
			"district":   district,
			"address":    address,
			"created_at": nowTime,
			"update_at":  nowTime,
		})
		if err != nil {
			return 0, err
		}
		addrId = addressId
	}
	return addrId, nil
}
