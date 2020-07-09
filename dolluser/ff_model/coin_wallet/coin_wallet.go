package coin_wallet

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type CoinWallet struct {
}

func NewCoinWallet() *CoinWallet {
	return &CoinWallet{}
}

func (p *CoinWallet) getTableName() string {
	return "coin_wallet"
}

func (p *CoinWallet) GetCoinWalletOneByUserId(userId int64, fields string) map[string]interface{} {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("user_id", userId).Limit(1).First()
	logrus.Debugf("Query coin_wallet by user_id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query coin_wallet by user_id. Error : %v", err)
		return nil
	}
	return one
}
