package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// md5解密

func Md5String(code string) string {
	md5 := md5.New()
	io.WriteString(md5, code)
	return hex.EncodeToString(md5.Sum(nil))
}
