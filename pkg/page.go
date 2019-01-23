package lib
type Page struct {
	pageNum     int
	pageSize   int
	totalPage  int
	totalCount int
	nextPage  bool
	list       interface{}
}
func PageUtil(count int, pageNum int, pageSize int, list interface{}) interface{} {
	nextPage := true
	tp := count / pageSize
	if count % pageSize > 0 {
		tp = count / pageSize + 1
		if tp <= pageNum {
			nextPage = false
		}
	}
	data := make(map[string]interface{})
	data["pageNum"] = pageNum
	data["pageSize"] = pageSize
	data["totalPage"] = tp
	data["totalCount"] = count
	data["nextPage"] = nextPage
	data["list"] = list
	return data
}
/**
获取偏移量
*/
func PageInit(num int,pageSize int)(offset int) {
	pageOffset := (num - 1) * pageSize
	return pageOffset
}