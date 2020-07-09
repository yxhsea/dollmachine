package ff_password

import (
	"bytes"
	"dollmachine/dollrpc/ff_core/ff_common/ff_crypt"
	"dollmachine/dollrpc/ff_core/ff_common/ff_random"
	"time"
)

func GenerateSalt() string {
	ranStr := time.Now().String() + ff_random.KrandAll(32)

	var cipherTxt bytes.Buffer
	cipherTxt.WriteString(ranStr)
	return ff_crypt.SHA1(cipherTxt.Bytes())
}

func GeneratePassword(password string, salt string) string {
	var str = "pw:" + password + ":salt:" + salt + ":wow:kael"

	var cipherTxt bytes.Buffer
	cipherTxt.WriteString(str)
	return ff_crypt.SHA1(cipherTxt.Bytes())
}

func Check(dbPassword string, dbSalt string, password string) bool {
	return dbPassword == GeneratePassword(password, dbSalt)
}
