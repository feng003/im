package tool

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
	"im/common/xerr"
)

// CopyFields 复制结构体字段
func CopyFields(dst, src interface{}) {
	dstValue := reflect.ValueOf(dst).Elem()
	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < dstValue.NumField(); i++ {
		dstField := dstValue.Field(i)
		srcField := srcValue.FieldByName(dstValue.Type().Field(i).Name)
		// 判断类型是否相同
		if dstField.Type() == srcField.Type() { // 判断类型是否相同
			// 相同类型进行赋值
			if srcField.IsValid() && dstField.CanSet() {
				dstField.Set(srcField)
			}
		}

	}
}

// CopyFieldsExcludeFields 复制结构体字段 - 排除指定字段
func CopyFieldsExcludeFields(dest, src interface{}, excludeFields []string) error {
	srcVal := reflect.ValueOf(src)
	destVal := reflect.ValueOf(dest)

	// 检查src是否为结构体
	if srcVal.Kind() != reflect.Struct {
		return fmt.Errorf("src must be a struct")
	}

	// 检查dest是否为指向结构体的指针
	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dest must be a pointer to a struct")
	}

	destElem := destVal.Elem()
	srcType := srcVal.Type()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcType.Field(i)
		exclude := false
		for _, fieldName := range excludeFields {
			if srcField.Name == fieldName {
				exclude = true
				break
			}
		}
		if !exclude {
			destField := destElem.FieldByName(srcField.Name)
			if destField.IsValid() && destField.CanSet() {
				destField.Set(srcVal.Field(i))
			}
		}
	}

	return nil
}

// ToJson 转换为json
func ToJson(d any) string {
	j, err := json.Marshal(d)
	if err != nil {
		log.Println("json Error : ", err)
		return ""
	}
	return string(j)
}

// CleanAllSpace 去掉字符串中的空格
func CleanAllSpace(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

// JSONStringToMap json转换为map
func JSONStringToMap(jsonString string) (map[string]string, error) {
	stringMap := map[string]string{}
	// 清除空格
	jsonString = strings.TrimSpace(jsonString)
	// 是否是空字符串
	if jsonString == "" {
		return stringMap, nil
	}

	// 定义一个map[string]interface{}类型的变量
	var data map[string]interface{}

	// 将JSON字符串解析到map中
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, err
	}

	// 将map[string]interface{}转换为map[string]string
	for key, value := range data {
		if strValue, ok := value.(string); ok {
			stringMap[key] = strValue
		} else {
			return nil, fmt.Errorf("Key %q 对应的值不是字符串类型", key)
		}
	}

	return stringMap, nil
}

// MultiParamsIsGt0 多参数判断大于0
func MultiParamsIsGt0[T int64 | float64](params ...T) bool {
	for _, param := range params {
		if param <= 0 {
			return false
		}
	}
	return true
}

// MultiParamsIsEmpty 多参数判空 - 只要有参数为空就返回true
func MultiParamsIsEmpty(params ...string) bool {
	for _, param := range params {
		if param == "" {
			return true
		}
	}
	return false
}

// ParamsNotIn 参数NotIn判断
func ParamsNotIn[T comparable](param T, paramsList []T) bool {
	return !ParamsIn(param, paramsList)
}

// ParamsIn 参数In判断
func ParamsIn[T comparable](param T, paramsList []T) bool {
	if len(paramsList) <= 0 {
		return false
	}
	has := false
	for _, v := range paramsList {
		if param == v {
			has = true
			break
		}
	}
	return has
}

// TrimStringFields 去除结构体中所有 string 类型字段的前后空格
func TrimStringFields(s interface{}) {
	v := reflect.ValueOf(s).Elem() // 获取结构体的值

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String {
			trimmed := strings.TrimSpace(field.String())
			field.SetString(trimmed)
		}
	}
}

// SliceSegmentByCharacter 切片按字符分割
func SliceSegmentByCharacter[T any](ids []T, character string) string {
	if len(ids) == 0 {
		return ""
	}

	var builder strings.Builder

	for i, id := range ids {
		if i > 0 {
			builder.WriteString(character)
		}
		builder.WriteString(fmt.Sprintf("%v", id))
	}

	return builder.String()
}

// SliceSegmentString 切片按字符分割字符串带引号
func SliceSegmentString(list []string) string {
	if len(list) == 0 {
		return ""
	}
	var builder strings.Builder

	for i, str := range list {
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(fmt.Sprintf("'%v'", str))
	}

	return builder.String()
}

// SliceUnique Slice去重
func SliceUnique[T comparable](slice []T) []T {
	var result []T
	m := make(map[T]bool)
	for _, item := range slice {
		if _, ok := m[item]; !ok {
			m[item] = true
			result = append(result, item)
		}
	}
	return result
}

// SliceStringToInt64 slice string转int64
func SliceStringToInt64(slice []string) (int64Slice []int64, err error) {
	for _, item := range slice {
		val, err := StringToInt64(item)
		if err != nil {
			return nil, err
		}
		int64Slice = append(int64Slice, val)
	}
	return int64Slice, nil
}

// StringToFloat64Slice 字符串切割为 []float64
func StringToFloat64Slice(s string) ([]float64, error) {
	slice := strings.Split(s, ",")
	float64Slice, err := SliceStringToFloat64(slice)
	if err != nil {
		return nil, err
	}
	return float64Slice, nil
}

// StringToInt64Slice 字符串切割为 []int64
func StringToInt64Slice(s string) ([]int64, error) {
	slice := strings.Split(s, ",")
	int64Slice, err := SliceStringToInt64(slice)
	if err != nil {
		return nil, err
	}
	return int64Slice, nil
}

// SliceStringToFloat64 slice string转float64
func SliceStringToFloat64(slice []string) (float64Slice []float64, err error) {
	for _, item := range slice {
		val, err := StringToFloat64(item)
		if err != nil {
			return nil, err
		}
		float64Slice = append(float64Slice, val)
	}
	return float64Slice, nil
}

// HideString 隐藏字符串替换为字符
func HideString(s string, front int, behind int, character string) string {
	runes := []rune(s)
	length := len(runes)
	if length <= front+behind {
		return s
	}
	frontStr := string(runes[:front])
	behindStr := string(runes[length-behind:])
	repeat := length - front - behind
	str := frontStr + strings.Repeat(character, repeat) + behindStr
	return str
}

// SetIdMap 设置id与数据的映射
func SetIdMap[T any](items []T) (map[int64]T, error) {
	result := make(map[int64]T)
	for _, item := range items {
		val := reflect.ValueOf(item)
		// If item is a pointer, get the value it points to
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		// Check if we have a valid struct
		if val.Kind() != reflect.Struct {
			return nil, fmt.Errorf("item is not a struct or pointer to struct")
		}
		idField := val.FieldByName("Id")
		if !idField.IsValid() || idField.Kind() != reflect.Int64 {
			return nil, fmt.Errorf("item does not have an int64 Id field")
		}
		id := idField.Int()
		result[id] = item
	}
	return result, nil
}

// GenerateUniqueString 生成指定长度的唯一字符串 - 长度越长越唯一
func GenerateUniqueString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "err: " + err.Error()
		}
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

// Param 参数结构
type Param struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// ParseParams 反解析查询参数
func ParseParams(params string) (map[string]Param, error) {
	params = strings.TrimSpace(params)
	paramsMap := map[string]Param{}
	// 判断有没有
	if params == "" {
		// 没有 - 返回
		return paramsMap, nil
	}
	// 有 - 解析
	err := json.Unmarshal([]byte(params), &paramsMap)
	if err != nil {
		return nil, errors.Errorf("Params参数解析失败！err : %v", err)
	}
	return paramsMap, nil
}

// ExcludeEqConditions 排除Eq类型的条件值
func ExcludeEqConditions(conditions []sq.Sqlizer, field string) []sq.Sqlizer {
	var result []sq.Sqlizer
	for _, condition := range conditions {
		if eq, ok := condition.(sq.Eq); ok {
			m := map[string]any(eq)
			if _, has := m[field]; has {
				continue
			}
		}
		result = append(result, condition)
	}
	return result
}

// StructToMap 结构体转Map
func StructToMap(input any) map[string]any {
	result := make(map[string]any)

	// 获取输入的反射值
	val := reflect.ValueOf(input)

	// 如果是指针，则解引用，获取其指向的值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// 确保输入是一个结构体
	if val.Kind() != reflect.Struct {
		return result
	}

	typ := val.Type()

	// 遍历结构体的字段
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		// 只处理导出的字段 (即字段名以大写字母开头)
		if field.PkgPath == "" {
			result[field.Name] = fieldValue.Interface()
		}
	}

	return result
}

// StructToMapWithJsonKey 结构体转map，带json标签作为key
func StructToMapWithJsonKey(input any) map[string]any {
	result := make(map[string]any)

	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return result
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		if field.PkgPath == "" {
			tag := field.Tag.Get("json")
			if tag != "" {
				result[tag] = fieldValue.Interface()
			}
		}
	}

	return result
}

// CamelToSnake 将驼峰字符串转换为小写加下划线
func CamelToSnake(s string) string {
	var result strings.Builder

	for i, r := range s {
		// 如果是大写字母，并且不是第一个字母，前面加下划线
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteByte('_')
			}
			// 将大写字母转换为小写
			result.WriteRune(unicode.ToLower(r))
		} else {
			// 非大写字母直接添加
			result.WriteRune(r)
		}
	}

	return result.String()
}

// MultiUpdateWhere 批量更新条件
type MultiUpdateWhere map[string]any

// MultiUpdateData 批量更新数据
type MultiUpdateData map[string]map[any]any

// MultiUpdateDataProcess 批量更新数据处理
func MultiUpdateDataProcess[T any](dataList []T) (multiWhere MultiUpdateWhere, multiData MultiUpdateData) {
	multiWhere = MultiUpdateWhere{}
	multiData = MultiUpdateData{}
	var ids []int64
	// 遍历数据
	var id int64
	for _, data := range dataList {
		// 数据结构体转map
		dataMap := StructToMap(data)
		// 取id
		id = dataMap["Id"].(int64)
		ids = append(ids, id)
		// 映射所有字段
		for k, v := range dataMap {
			// id字段则跳过
			if k == "Id" {
				continue
			}

			// 设置字段key和值
			key := CamelToSnake(k)
			value := v
			// 自动处理更新时间
			if key == "update_time" || key == "update_at" {
				value = GetCurrentUnix()
			}

			// 设置数据 - 判断有没有
			if _, ok := multiData[key]; !ok {
				multiData[key] = map[any]any{
					id: value,
				}
			} else {
				multiData[key][id] = value
			}
		}

	}
	// 设置条件
	multiWhere["Id"] = ids
	return
}

// CopyDataList 复制数据列表
func CopyDataList[S, T any](sourceList []*S, targetList []*T, tag string) ([]*T, error) {
	for _, v := range sourceList {
		var data T
		err := copier.Copy(&data, &v)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.COPY_DATA_ERROR), "[Tag: %s] copy data error %v \n", tag, err)
		}
		targetList = append(targetList, &data)
	}
	return targetList, nil
}

// IfThen 三元运算
func IfThen[T any](condition bool, trueValue T, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

// StructHasTag 检测结构体是否包含指定tag字段
func StructHasTag(structValue interface{}, fieldFlag, fieldName string) bool {
	rv := reflect.ValueOf(structValue)
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		tagName := field.Tag.Get(fieldFlag)
		if tagName == fieldName {
			return true
		}
	}
	return false
}

// MatchName 检查姓名是否与带通配符的尾号匹配
// 示例 :
// "王芳" -> "*王芳" true
// "张三" -> "*王芳" false
func MatchName(fullName, pattern string) bool {
	// 去掉 pattern 中的通配符（*），保留需要匹配的尾号
	patternSuffix := strings.TrimLeft(pattern, "*")

	// 如果尾号长度超过 fullName 长度，则无法匹配
	if len(patternSuffix) > len(fullName) {
		return false
	}

	// 获取 fullName 的尾部与 patternSuffix 比较
	fullNameSuffix := fullName[len(fullName)-len(patternSuffix):]
	return fullNameSuffix == patternSuffix
}

// HidePhone 隐藏手机号码【中间四位】
func HidePhone(phone string) string {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return ""
	}
	// 判断是否包含脱敏字符 *
	if strings.Contains(phone, "*") {
		// 已脱敏 - 无需重复处理
		return phone
	}
	if len(phone) != 11 {
		return "Invalid phone number"
	}
	return phone[:3] + "****" + phone[7:]
}

// HideUsername 对用户名姓名
// 隐藏规则：1个字显示，2个字显示姓，3个字隐藏中间，大于3个字只显示首尾
func HideUsername(username string) string {
	username = strings.TrimSpace(username)
	// 判断是否包含脱敏字符 *
	if strings.Contains(username, "*") {
		// 已脱敏 - 无需重复处理
		return username
	}

	// 获取用户名的长度
	length := utf8.RuneCountInString(username)
	if length == 0 {
		return username // 如果用户名为空，直接返回
	}

	// 将字符串转换为 rune 切片，方便处理多字节字符
	runes := []rune(username)

	switch {
	case length == 1:
		// 1个字：直接显示
		return string(runes[0])
	case length == 2:
		// 2个字：只显示姓，第二个字用 * 替换
		return string(runes[0]) + "*"
	case length == 3:
		// 3个字：显示第一个字和最后一个字，中间用 * 替换
		return string(runes[0]) + "*" + string(runes[2])
	default:
		// 超过3个字：显示第一个字和最后一个字，中间用 * 替换
		firstChar := string(runes[0])
		lastChar := string(runes[length-1])
		maskedPart := firstChar
		for i := 1; i < length-1; i++ {
			maskedPart += "*"
		}
		maskedPart += lastChar
		return maskedPart
	}
}

// HideIDCard 隐藏身份证
func HideIDCard(idCard string) string {
	idCard = strings.TrimSpace(idCard)
	if idCard == "" {
		return ""
	}
	// 判断是否包含脱敏字符 *
	if strings.Contains(idCard, "*") {
		// 已脱敏 - 无需重复处理
		return idCard
	}
	return HideString(idCard, 4, 4, "*")
}

// FloatToPercent 小数转百分比
func FloatToPercent(f float64) string {
	return fmt.Sprintf("%.2f%%", f*100)
}

// SliceRemoveSliceElements 从切片中移除多个元素
func SliceRemoveSliceElements[T comparable](slice []T, elements []T) []T {
	result := make([]T, 0)
	for _, v := range elements {
		result = SliceRemoveElement(slice, v)
	}
	return result
}

// SliceRemoveElement 从切片中移除指定的元素
func SliceRemoveElement[T comparable](slice []T, element T) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}

// CompareErrorCode 比较错误码是否相等
func CompareErrorCode(code int, err error) bool {
	errCode := ConvertErrorCode(err)
	return code == errCode
}

// ConvertErrorCode 转换错误Code
func ConvertErrorCode(err error) int {
	if err == nil {
		return 0
	}
	// 断言
	causeErr := errors.Cause(err)
	var e *xerr.CodeError
	if errors.As(causeErr, &e) {
		code := e.GetErrCode()
		return int(code)
	}
	return 0
}
