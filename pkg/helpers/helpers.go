// Package helpers 存放辅助方法
package helpers

import (
	"crypto/rand"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"time"

	mathrand "math/rand"

	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// Empty 类似于 PHP 的 empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

// RandomNumber 生成长度为 length 随机数字字符串
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// FirstElement 安全地获取 args[0]，避免 panic: runtime error: index out of range
func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

// RandomString 生成长度为 length 的随机字符串
func RandomString(length int) string {
	mathrand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}
	return string(b)
}

func UUID() string {
	var uuidV1 uuid.UUID
	var err error

	for i := 0; i < 9999; i++ {
		uuidV1, err = uuid.NewUUID()
		// time.Sleep(1000 * time.Nanosecond)
		if err != nil {
			continue
		} else {
			break
		}
	}
	return cast.ToString(uuidV1)
}

// IsUUID checks if a string is a valid UUID.
func IsUUIDV1(u string) bool {
	// Regular expression to match a valid UUID
	r := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
	return r.MatchString(u)
}

func IsUUID(u string) bool {
	err := uuid.Validate(u)
	if err == nil {
		return true
	}
	return false
}

// input: "2024-05-27T15:12:48.356854+08:00"
// output: "2024-05-27T07:12:48.356854Z"
func ToUTC(input string) string {
	// 解析输入时间字符串
	t, err := time.Parse(time.RFC3339Nano, input)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		t = time.Now().UTC()
	}
	// 转换为 UTC 时间
	utcTime := t.UTC()

	// 格式化为 ISO 8601 格式的字符串
	output := utcTime.Format(time.RFC3339Nano)

	return output
}
