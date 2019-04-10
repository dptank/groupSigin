package ex
var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_ADD_FAIL: 			 	 "保存失败",
	ERROR_ID_FAIL:					 "id不能为空！",
}
/**
获取错误信息
*/
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
