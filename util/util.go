package util

import (
	"math/rand"
	"strconv"
	"time"
)

// CreateId id生成器
func CreateId() int64 {
	rand.Seed(time.Now().Unix())
	//当前时间戳
	str1 := strconv.Itoa(int(time.Now().Unix()))
	//随机4位数字
	str2 := RandomString(4)
	id, _ := strconv.ParseInt(str1+str2, 10, 64)
	return id
}

//RandomString 生成随机数字字符串
func RandomString(len int) string {
	bytes := make([]byte, len)
	rand.Seed(time.Now().Unix())
	for i := 0; i < len; i++ {
		bytes[i] = byte(49 + rand.Intn(9))
	}
	return string(bytes)
}
