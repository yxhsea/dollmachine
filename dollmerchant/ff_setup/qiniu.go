package ff_setup

import (
	"github.com/qiniu/api.v7/storage"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

func SetupQiNiu(accessKey string, secretKey string, zone string) error {
	ff_vars.QiNiuAccessKey = accessKey
	ff_vars.QiNiuSecretKey = secretKey
	//ff_vars.QiNiuZone = zone
	switch zone {
	case "huanan":
		ff_vars.QiNiuZone = &storage.ZoneHuanan
		break
	case "huadong":
		ff_vars.QiNiuZone = &storage.ZoneHuadong
		break
	case "huabei":
		ff_vars.QiNiuZone = &storage.ZoneHuabei
		break
	case "beimei":
		ff_vars.QiNiuZone = &storage.ZoneBeimei
		break
	}

	return nil
}
