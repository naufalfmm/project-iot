package consts

const (
	ResponseCode = "response_code"
	PostgreTrx   = "trx"
	UserLoginKey = "user_login"
)

const AsciiCodeSpace = 32

var (
	UnitSensorStandard map[string]string = map[string]string{
		"temp":     "°C",
		"tds":      "ppm",
		"humidity": "%",
	}
)
