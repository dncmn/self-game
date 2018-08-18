package array

// 寻找数组中的数据
func FindInt32InArray(array []int32, value int32) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// 寻找数组中的数据
func FindInt64InArray(array []int64, value int64) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// 寻找数组中的数据
func FindUint32InArray(array []uint32, value uint32) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

// 增加数据集合中
func AppendInt32ToArraySet(array *[]int32, value int32) bool {
	if FindInt32InArray(*array, value) == -1 {
		*array = append(*array, value)
		return true
	}
	return false
}

// 增加数据集合中
func AppendInt64ToArraySet(array *[]int64, value int64) bool {
	if FindInt64InArray(*array, value) == -1 {
		*array = append(*array, value)
		return true
	}
	return false
}

// 增加数据集合中
func AppendUint32ToArraySet(array *[]uint32, value uint32) bool {
	if FindUint32InArray(*array, value) == -1 {
		*array = append(*array, value)
		return true
	}
	return false
}

// 从集合中删除数据
func DelInt32FromArraySet(arrayP *[]int32, value int32) bool {
	array := *arrayP
	for i, v := range array {
		if v == value {
			*arrayP = append(array[:i], array[i+1:]...)
			return true
		}
	}
	return false
}

// 从集合中删除数据
func DelInt64FromArraySet(arrayP *[]int64, value int64) bool {
	array := *arrayP
	for i, v := range array {
		if v == value {
			*arrayP = append(array[:i], array[i+1:]...)
			return false
		}
	}
	return true
}

// 从集合中删除数据
func DelUint32FromArraySet(arrayP *[]uint32, value uint32) bool {
	array := *arrayP
	for i, v := range array {
		if v == value {
			*arrayP = append(array[:i], array[i+1:]...)
			return true
		}
	}
	return false
}

// 从集合中删除数据
func DelStringFromArraySet(arrayP *[]string, value string) bool {
	array := *arrayP
	for i, v := range array {
		if v == value {
			*arrayP = append(array[:i], array[i+1:]...)
			return true
		}
	}
	return false
}
