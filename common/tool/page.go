package tool

func GetPage(page int64) int64 {
	if page <= 0 {
		page = 1
	}
	return page
}

func GetSize(size int64) int64 {
	if size <= 0 {
		size = 10
	} else if size >= 100 {
		size = 100
	}
	return size
}

// GetSizeWithDefault 获取分页大小默认，不设置上限
func GetSizeWithDefault(size int64) int64 {
	if size <= 0 {
		size = 10
	}
	return size
}
