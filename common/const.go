package common

const (
	// STATUS_OK 成功
	STATUS_OK      = 0    // 成功
	STATUS_NO_DATA = 2    // 无数据
	ERROR_INTER    = 1000 // 参数错误
	ERROR_PARAM    = 1001 // 服务异常|内部错误
	ERROR_CAPTCHA  = 1002 // 验证码错误
)

var INFO = map[int]string{
	ERROR_PARAM:   "参数错误",
	ERROR_INTER:   "内部错误",
	ERROR_CAPTCHA: "验证码错误",
}

const (
	TOKEN  = "bHzz89fXlVBJZllDpVgeHw7ymmMLoHU"
	AESKEY = "90yTvpl6UckAl3BFXBg0RvPFH73kOuFstr7ZipMhG2i"
	CORPID = "wwff0fd981753d18a7"
)
