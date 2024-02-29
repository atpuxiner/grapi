package status

const (
	CodeSuccess = 0
	CodeFailure = 1

	CodeBadParam       = 40000
	CodeUnauthorized   = 40001
	CodeRecordNotFound = 40004
	CodeRecordExisted  = 40005

	CodeServerError = 50000
)

var CodeMap map[int]string

func CodeMsg(code int) string {
	if msg, ok := CodeMap[code]; ok {
		return msg
	}
	return ""
}

func init() {
	CodeMap = map[int]string{
		CodeSuccess:        "请求成功",
		CodeFailure:        "请求失败",
		CodeBadParam:       "参数错误",
		CodeUnauthorized:   "认证失败",
		CodeRecordNotFound: "记录未找到",
		CodeRecordExisted:  "记录已存在",
		CodeServerError:    "服务器错误",
	}
}
