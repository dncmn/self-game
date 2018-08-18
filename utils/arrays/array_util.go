package arrays

import "sort"

// ReverseInts:数组降序排序
func ReverseInts(array []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	return array
}

// SortedInts:数组正序排序
func SortedInts(array []int) []int {
	sort.Ints(array)
	return array
}

// FindIndexInArray:找出数组中,元素的位置
func FindIndexInArray(a []int, x int) int {
	sort.Ints(a)
	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		return i
	} else {
		i = -1
	}
	return i
}

// 找出int数组中不同的元素
func FindDiffEles(max, min []int) (result []int) {
	// 将数量较少的数组添加到map中
	mins := make(map[int]bool, len(min))
	for _, val := range min {
		mins[val] = true
	}

	// 从最大的数组中找出在这个map中不存在的元素,并且添加到返回值中
	for _, val := range max {
		if _, ok := mins[val]; !ok {
			result = append(result, val)
		}
	}
	return result
	return
}
