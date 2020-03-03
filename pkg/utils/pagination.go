package utils

func GetOffset(pageNum uint,pageSize uint) (uint,uint) {
	// 处理size
	var limit uint
	limit=1
	if pageSize > 1{
		limit=pageSize
	}
	// 处理offset
	var offset uint
	offset=0
	if pageNum > 0 {
		offset = (pageNum - 1) * pageSize
	}

	return offset,limit
}
