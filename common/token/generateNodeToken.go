package token

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func GenerateNodeToken(label string) string {
	hasher := md5.New()
	unixNow := time.Now().Unix()
	unixNowStr := strconv.FormatInt(unixNow, 10)

	txt := fmt.Sprintf("%s%s", label, unixNowStr)

	hash := hasher.Sum([]byte(txt))

	return hex.EncodeToString(hash[:])
}
