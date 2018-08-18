package conv

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const (
	_BIT_CHINESE_CODE0 = byte(0x80)
	_BIT_CHINESE_CODE1 = byte(0xE0)
	_BIT_CHINESE_CODE2 = byte(0x80)
	_BIT_CHINESE_CODE3 = byte(0x80)
)

func Str2Int64(str string) (int64, error) {
	v, err := strconv.ParseInt(str, 10, 64)
	return v, err
}

func Str2Int64s(str string, splitStr string) ([]int64, error) {
	if str == "" {
		return nil, fmt.Errorf("Empty string")
	}
	splittedStr := strings.Split(str, splitStr)
	result := make([]int64, len(splittedStr))
	for i, intStr := range splittedStr {
		value, err := Str2Int64(intStr)
		if err != nil {
			return nil, err
		}
		result[i] = value
	}
	return result, nil
}

func Str2Uint64(str string) (uint64, error) {
	v, err := strconv.ParseUint(str, 10, 64)
	return v, err
}

func Str2Int32(str string) (int32, error) {
	v, err := strconv.ParseInt(str, 10, 32)
	return int32(v), err
}

func Str2Int32s(str string, splitStr string) ([]int32, error) {
	if str == "" {
		return nil, fmt.Errorf("Empty string")
	}
	splittedStr := strings.Split(str, splitStr)
	result := make([]int32, len(splittedStr))
	for i, intStr := range splittedStr {
		value, err := Str2Int32(intStr)
		if err != nil {
			return nil, err
		}
		result[i] = value
	}
	return result, nil
}

func Str2Byte(str string) (byte, error) {
	v, err := strconv.ParseInt(str, 10, 8)
	return byte(v), err
}

func Str2Bool(str string) (bool, error) {
	v, err := strconv.ParseInt(str, 10, 8)
	if err != nil || v != 1 {
		return false, err
	}
	return true, nil
}

func Str2Uint8s(str string) ([]byte, error) {
	if str == "" {
		return nil, fmt.Errorf("Empty string")
	}
	splittedStr := strings.Split(str, "-")
	result := make([]byte, len(splittedStr))
	for i, intStr := range splittedStr {
		value, err := Str2Byte(intStr)
		if err != nil {
			return nil, err
		}
		result[i] = value
	}
	return result, nil
}

func Str2Uint32(str string) (uint32, error) {
	v, err := strconv.ParseUint(str, 10, 32)
	return uint32(v), err
}

func Str2Int(str string) (int, error) {
	v, err := strconv.ParseInt(str, 10, 32)
	return int(v), err
}

func Str2Float32(str string) (float32, error) {
	f32, err := strconv.ParseFloat(str, 32)
	return float32(f32), err
}

func Float322Str(v float32) string {
	return strconv.FormatFloat(float64(v), 'G', 5, 32)
}

func Float642Str(v float64) string {
	return strconv.FormatFloat(v, 'G', 5, 64)
}

func Int32Str(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

func IntStr(v int) string {
	return strconv.FormatInt(int64(v), 10)
}

func Uint32Str(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

func Int64Str(v int64) string {
	return strconv.FormatInt(v, 10)
}

func Uint64Str(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// 判断UTF-8字符串是否全是汉字和英文以及标点符号
func IsAllGB2312(str string) bool {
	for _, r := range str {
		if !(unicode.IsNumber(r) ||
			unicode.IsLetter(r) ||
			unicode.Is(unicode.Scripts["Han"], r)) {
			return false
		}
	}
	// for _, r := range str {
	// 	if !(unicode.Is(unicode.Scripts["Han"], r) ||
	// 		unicode.Is(unicode.Scripts["Katakana"], r) ||
	// 		unicode.Is(unicode.Scripts["Hiragana"], r) ||
	// 		unicode.IsLetter(r) ||
	// 		unicode.IsDigit(r)) ||
	// 		unicode.IsControl(r) ||
	// 		unicode.IsPunct(r) ||
	// 		unicode.IsSpace(r) {
	// 		return false
	// 	}
	// }
	return true
}

// check name length
func NameLength(str string) int {
	var length int
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			// unicode.Is(unicode.Scripts["Katakana"], r) ||
			// unicode.Is(unicode.Scripts["Hiragana"], r) {
			length += 2
		} else {
			length += 1
		}
	}
	return length
}

// 判断字符中是否含有空格
func IsHasSpace(str string) bool {
	for _, r := range str {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

func IsLegalName(str string) bool {
	for _, r := range str {
		if unicode.IsSpace(r) {
			return false
		}
		if !unicode.IsGraphic(r) {
			return false
		}
		if unicode.IsPunct(r) {
			return false
		}
		if unicode.IsSymbol(r) {
			return false
		}
	}
	return true
}

func StringsToInt(strs []string) []int {
	l := len(strs)
	ret := make([]int, l)
	for k, _ := range strs {
		v, _ := strconv.ParseInt(strs[k], 10, 32)
		ret[k] = int(v)
	}
	return ret
}

func StringsToInt64(strs []string) []int64 {
	l := len(strs)
	ret := make([]int64, l)
	for k, _ := range strs {
		v, _ := strconv.ParseInt(strs[k], 10, 64)
		ret[k] = v
	}
	return ret
}

func ValsToInt64(strs []string) ([]string, []int64) {
	l := len(strs)
	ret1, ret2 := make([]string, l/2), make([]int64, l/2)
	n := 0
	for k := 0; k < l; k += 2 {
		ret1[n] = strs[k]
		v2, _ := strconv.ParseInt(strs[k+1], 10, 64)
		ret2[n] = v2
		n++
	}
	return ret1, ret2
}

func StringsToInt32(strs []string) []int32 {
	l := len(strs)
	ret := make([]int32, l, l)
	for k, _ := range strs {
		v, _ := strconv.ParseInt(strs[k], 10, 32)
		ret[k] = int32(v)
	}
	return ret
}

func StringsToUint64(strs []string) []uint64 {
	l := len(strs)
	ret := make([]uint64, l)
	for k, _ := range strs {
		v, _ := strconv.ParseUint(strs[k], 10, 64)
		ret[k] = uint64(v)
	}
	return ret
}

func ValsToUint64(strs []string) ([]string, []uint64) {
	l := len(strs)
	ret1, ret2 := make([]string, l/2), make([]uint64, l/2)
	n := 0
	for k := 0; k < l; k += 2 {
		ret1[n] = strs[k]
		v2, _ := strconv.ParseUint(strs[k+1], 10, 64)
		ret2[n] = uint64(v2)
		n++
	}
	return ret1, ret2
}

func StrsToMapInt32(strs []string) map[string]int32 {
	l := len(strs)
	ret := make(map[string]int32, l)
	if l > 0 {
		for i := 0; i < l; i += 2 {
			v, _ := strconv.ParseInt(strs[i+1], 10, 32)
			ret[strs[i]] = int32(v)
		}
	}
	return ret
}

func StrsToMapInt64String(strs []string) map[int64]string {
	l := len(strs)
	ret := make(map[int64]string, l)
	if l > 0 {
		for i := 0; i < l; i += 2 {
			v, _ := strconv.ParseInt(strs[i], 10, 64)
			ret[v] = strs[i+1]
		}
	}
	return ret
}

func Int32sToString(int32s []int32, sep string) string {
	if len(int32s) == 0 {
		return ""
	}
	if len(int32s) == 1 {
		return strconv.Itoa(int(int32s[0]))
	}
	a := make([]string, len(int32s))
	for i := 0; i < len(int32s); i++ {
		a[i] = strconv.Itoa(int(int32s[i]))
	}

	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func Int32Map2StrMap(input map[int32]int32) (output map[string]int32) {
	output = make(map[string]int32)
	for k, v := range input {
		output[Int32Str(k)] = v
	}
	return
}

func SortMapByStringKey(mapValues map[string]string) []string {
	keys := make([]int, len(mapValues))
	i := 0
	for k, _ := range mapValues {
		v, _ := Str2Int(k)
		keys[i] = v
		i++
	}
	sort.Ints(keys)
	result := make([]string, i)
	for i = 0; i < len(result); i++ {
		k := IntStr(keys[i])
		result[i] = mapValues[k]
	}
	return result
}

func SortMapByIntKey(mapValues map[int]string) []string {
	keys := make([]int, len(mapValues))
	i := 0
	for k, _ := range mapValues {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	result := make([]string, i)
	for i = 0; i < len(result); i++ {
		result[i] = mapValues[keys[i]]
	}
	return result
}

type int64Slice []int64

func (p int64Slice) Len() int           { return len(p) }
func (p int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p int64Slice) Sort()              { sort.Sort(p) }

func SortMapByInt64Key(mapValues map[int64]string) []string {
	keys := make([]int64, len(mapValues))
	i := 0
	for k, _ := range mapValues {
		keys[i] = k
		i++
	}
	int64Slice(keys).Sort()
	result := make([]string, i)
	for i = 0; i < len(result); i++ {
		result[i] = mapValues[keys[i]]
	}
	return result
}

func SortInt64(values []int64) {
	int64Slice(values).Sort()
}

func SliceStringsToMap(str []string) map[string]string {
	size := len(str)
	if size > 0 {
		ret := make(map[string]string, size)
		for i := 0; i < size; i += 2 {
			ret[string(str[i])] = string(str[i+1])

		}
		return ret
	}
	return map[string]string{}
}

func URLValueToMap(v url.Values) map[string]string {
	ret := make(map[string]string)
	for k, vs := range v {
		ret[k] = vs[0]
	}
	return ret
}

// limit to 0-10 show star number
func Int32ToLocalNumber(v int32, lang string) string {
	var (
		chineseNumbers = [...]string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
		arabicNumbers  = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	)
	if v >= 0 && v <= 10 {
		switch lang {
		case "zh_CN", "zh_TW":
			return chineseNumbers[v]
		default:
			return arabicNumbers[v]
		}
	}
	return ""
}

// limit to 0-10
func Int32ToChinese(v int32) string {
	var chineseNumbers = []string{
		"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十",
	}
	if v >= 0 && v <= 10 {
		return chineseNumbers[v]
	}
	return ""
}
