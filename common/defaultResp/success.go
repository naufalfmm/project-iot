package defaultResp

type Success struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func CreateSuccessResp(code int, data interface{}) Success {
	return Success{
		Code: code,
		Data: data,
	}
}
