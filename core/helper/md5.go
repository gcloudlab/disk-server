package helper

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
