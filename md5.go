package goutil

import (
    "crypto/md5"
    "fmt"
)

func GetMD5(body []byte) string{
    return fmt.Sprintf("%x",md5.Sum(body))
}
