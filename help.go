package common

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

// ChangeType 变量类型转换
// val 变量值 ty 转换成类型
// 支持string,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64,bool类型
func ChangeType(val interface{}, ty string) interface{} {

	f := float64(0)
	s := ""
	isString := false
	switch value := val.(type) {
	case int:
		f = float64(value)
	case int8:
		f = float64(value)
	case int16:
		f = float64(value)
	case int32:
		f = float64(value)
	case int64:
		f = float64(value)
	case uint:
		f = float64(value)
	case uint8:
		f = float64(value)
	case uint16:
		f = float64(value)
	case uint32:
		f = float64(value)
	case uint64:
		f = float64(value)
	case float32:
		f = float64(value)
	case float64:
		f = value
	case string:
		s = value
		isString = true
	case bool:
		if value {
			f = float64(1)
		} else {
			f = float64(0)
		}
		s = strconv.FormatBool(value)
	}
	if "" == s {
		str := strings.Split(fmt.Sprintf("%v", val), ".")
		le := 0
		if len(str) > 1 {
			le = len(str[1])
		}
		s = big.NewRat(1, 1).SetFloat64(f).FloatString(le)
	}
	if "string" == ty {
		return s
	}
	if "bool" == ty {
		res, _ := strconv.ParseBool(s)
		return res
	}
	if isString {
		f, _ = strconv.ParseFloat(s, 64)
	}

	switch ty {
	case "int":
		return int(f)
	case "int8":
		return int8(f)
	case "int16":
		return int16(f)
	case "int32":
		return int32(f)
	case "int64":
		return int64(f)
	case "uint":
		return uint(f)
	case "uint8":
		return uint8(f)
	case "uint16":
		return uint16(f)
	case "uint32":
		return uint32(f)
	case "uint64":
		return uint64(f)
	case "float32":
		return float32(f)
	case "float64":
		return f
	case "string":
		return s
	case "bool":
		res, _ := strconv.ParseBool(s)
		return res
	}

	return val
}

// InArray 判断元素是否在数组中
// 支持string,int,int8,int16,int32,int64,float32,float64,bool类型
func InArray(item interface{}, arr interface{}) bool {
	switch arr := arr.(type) {
	case []string:
		for _, v := range arr {
			if v == item {
				return true
			}
		}
	case []int8:
		for _, v := range arr {
			if v == item {
				return true
			}
		}
	case []int16:
		for _, v := range arr {
			if v == item {
				return true
			}
		}
	case []int32:
		for _, v := range arr {
			if v == item {
				return true
			}
		}
	case []int:
		for _, v := range arr {
			if v == item {
				return true
			}
		}
	case []int64:
		for _, v := range arr {
			if v == item {
				return true
			}
		}
	case []float32:
		for _, v := range arr {
			if v == item {
				return true
			}
		}
	case []float64:
		for _, v := range arr {
			if v == item {
				return true
			}
		}
	case []bool:
		for _, v := range arr {
			if v == item {
				return true
			}
		}
	}
	return false
}

// GenTree 将数组map生成tree
// items 将要排序的数组
// addEmptyChild child为空时是否返回空数组，默认true
// option[0] id别名，默认id
// option[1] pid别名，默认pid
// option[2] child别名，默认返回child
// example:
//  items := []map[string]interface{}{
//	   {"id": "1", "name": "1"},
//	   {"id": "2", "pid": "1", "name": "2"},
//	   {"id": "3", "pid": "1", "name": "3"},
//	   {"id": "4", "pid": "2", "name": "3"},
//   }
// JsonEncode(GenTree(items, true))
func GenTree(items []map[string]interface{}, addEmptyChild bool, option ...string) []map[string]interface{} {
	// 赋默认值
	_option := []string{"id", "pid", "child"}
	if len(option) > 0 {
		for k, v := range option {
			_option[k] = v
		}
	}

	newData := make(map[interface{}]map[string]interface{})
	for k, v := range items {
		if addEmptyChild {
			items[k][_option[2]] = []interface{}{}
		}
		newData[v[_option[0]]] = v
	}

	var tree []map[string]interface{}
	for _, v := range newData {
		_, find := v[_option[1]]
		var b bool
		if find {
			if _, ok := newData[v[_option[1]]]; ok {
				b = true
			}
		}
		if b {
			newData[v[_option[1]]][_option[2]] = append(newData[v[_option[1]]][_option[2]].([]interface{}), newData[v[_option[0]]])
		} else {
			tree = append(tree, newData[v[_option[0]]])
		}
	}
	return tree
}

//MapKeys 获取map里面所有键名
func MapKeys(items map[string]interface{}) []string {
	i, keys := 0, make([]string, len(items))
	for key := range items {
		keys[i] = key
		i++
	}
	return keys
}

//loc, _ := time.LoadLocation("Asia/Shanghai")
//t, _ := time.Parse("2006-01-02 15:04:05", v.EventTime)
//dt := time.Unix(t.Unix(), 0).In(loc).Format("2006-01-02 15:04:05")
//v.EventTime = dt

// GetAndDefault 获取map里面的值，如果没有则返回默认值
func GetAndDefault(items map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if v, ok := items[key]; ok {
		return v
	}
	return defaultValue
}

// MapSort map排序
// arr 数组
// key 排序的键
// desc 是否降序
func MapSort(arr []map[string]interface{}, key string, desc bool) {
	sort.Slice(arr, func(i, j int) bool {
		ii := ChangeType(arr[i][key], "string").(string)
		jj := ChangeType(arr[j][key], "string").(string)
		if desc {
			return ii > jj
		}
		return ii < jj
	})
}

// Wordwrap 字符串换行
// str 字符串
// width 换行的宽度
// breakStr 换行的字符
func Wordwrap(str string, width int, breakStr string) string {
	if width == 0 {
		return str
	}
	if len(str) <= width {
		return str
	}
	var result string
	for i := 0; i < len(str); i += width {
		if i+width > len(str) {
			result += str[i:]
		} else {
			result += str[i : i+width]
		}
		result += breakStr
	}
	return result
}

func Mul(a, b float64, more ...float64) (float64, error) {
	bigF1 := new(big.Float).SetFloat64(a)

	bigF2 := new(big.Float).SetFloat64(b)

	pf := new(big.Float).Mul(bigF1, bigF2)

	if len(more) > 0 {
		for _, m := range more {
			floatM := new(big.Float).SetFloat64(m)
			pf = new(big.Float).Mul(pf, floatM)
		}
	}

	return strconv.ParseFloat(pf.String(), 64)
}

func MulCeil(a, b float64, more ...float64) (result float64, err error) {
	result, err = Mul(a, b, more...)
	if err != nil {
		return
	}
	result, err = FormatFloatCeil(result, 2)
	if err != nil {
		return
	}
	return
}

func Add(a, b float64, more ...float64) (float64, error) {
	bigF1 := new(big.Float).SetFloat64(a)

	bigF2 := new(big.Float).SetFloat64(b)

	pf := new(big.Float).Add(bigF1, bigF2)

	if len(more) > 0 {
		for _, m := range more {
			floatM := new(big.Float).SetFloat64(m)
			pf = new(big.Float).Add(pf, floatM)
		}
	}

	return strconv.ParseFloat(pf.String(), 64)
}

func AddCeil(a, b float64, more ...float64) (result float64, err error) {
	result, err = Add(a, b, more...)
	if err != nil {
		return
	}
	result, err = FormatFloatCeil(result, 2)
	if err != nil {
		return
	}
	return
}

func Sub(a, b float64, more ...float64) (float64, error) {
	bigF1 := new(big.Float).SetFloat64(a)

	bigF2 := new(big.Float).SetFloat64(b)

	pf := new(big.Float).Sub(bigF1, bigF2)

	if len(more) > 0 {
		for _, m := range more {
			floatM := new(big.Float).SetFloat64(m)
			pf = new(big.Float).Sub(pf, floatM)
		}
	}

	return strconv.ParseFloat(pf.String(), 64)
}

func SubCeil(a, b float64, more ...float64) (result float64, err error) {
	result, err = Sub(a, b, more...)
	if err != nil {
		return
	}
	result, err = FormatFloatCeil(result, 2)
	if err != nil {
		return
	}
	return
}

func Div(a, b float64, more ...float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}

	bigF1 := new(big.Float).SetFloat64(a)

	bigF2 := new(big.Float).SetFloat64(b)

	pf := new(big.Float).Quo(bigF1, bigF2)

	if len(more) > 0 {
		for _, m := range more {
			if m == 0 {
				return 0, errors.New("除数不能为0")
			}
			floatM := new(big.Float).SetFloat64(m)
			pf = new(big.Float).Quo(pf, floatM)
		}
	}

	return strconv.ParseFloat(pf.String(), 64)
}

func DivCeil(a, b float64, more ...float64) (result float64, err error) {
	result, err = Div(a, b, more...)
	if err != nil {
		return
	}
	result, err = FormatFloatCeil(result, 2)
	if err != nil {
		return
	}
	return
}

func FormatFloatCeil(num float64, decimal int) (float64, error) {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	price, err := Mul(num, d)
	if err != nil {
		return 0, nil
	}
	ceil := math.Ceil(price)
	price, err = Div(ceil, d)
	if err != nil {
		return 0, nil
	}
	res := strconv.FormatFloat(price, 'f', -1, 64)
	return strconv.ParseFloat(res, 64)
}
