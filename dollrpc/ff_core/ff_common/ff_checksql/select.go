package ff_checksql

import (
	"errors"
	"github.com/sirupsen/logrus"
	"dollmachine/dollrpc/ff_core/ff_common/ff_json"
)

func CheckSelectResult(args interface{}, affectRow int64, err error, mustHasRow bool, msg string, targetId interface{}) error {
	if err != nil {
		fields := make(logrus.Fields)
		fields["ff_checksql"] = "CheckSelectResult"
		fields["targetId"] = targetId
		fields["err"] = err.Error()
		fields["args"] = ff_json.MarshalToStringNoError(args)
		logrus.WithFields(fields).Error(msg)
		return err
	}
	if mustHasRow && affectRow <= 0 {
		fields := make(logrus.Fields)
		fields["ff_checksql"] = "CheckSelectResult"
		fields["targetId"] = targetId
		fields["affectRow"] = "empty"
		fields["args"] = ff_json.MarshalToStringNoError(args)
		logrus.WithFields(fields).Error(msg)
		return errors.New("affectRow empty")
	}
	return nil
}
