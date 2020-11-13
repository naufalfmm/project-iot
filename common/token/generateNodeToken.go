package token

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"time"
)

func GenerateNodeToken(label string) string {
	unixNow := time.Now().Unix()
	unixNowStr := strconv.FormatInt(unixNow, 10)

	h := hmac.New(sha1.New, []byte(unixNowStr))
	h.Write([]byte(label))
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
