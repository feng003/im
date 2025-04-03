package tool

import (
	"crypto/md5"
	"fmt"
)

func Md5ByString(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return fmt.Sprintf("%x", m.Sum(nil))
}

func Md5ByStringWithSalt(str, salt string) string {
	saltedStr := str + salt
	hash := md5.Sum([]byte(saltedStr))
	return fmt.Sprintf("%x", hash)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func Md5ByBytesWithSalt(b []byte, salt string) string {
	saltedBytes := append(b, []byte(salt)...)
	return fmt.Sprintf("%x", md5.Sum(saltedBytes))
}
