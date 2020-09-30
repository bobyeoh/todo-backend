package utils

import (
	"crypto/md5"
	"fmt"
)

// MD5 godoc
func MD5(str string) string {
	data := []byte("+todo+" + str + "+todo+")
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
