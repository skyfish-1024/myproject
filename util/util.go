package util

import (
	"math/rand"
	"strconv"
	"time"
)

func CreateId() int64 {
	rand.Seed(time.Now().Unix())

	str1 := strconv.Itoa(int(time.Now().Unix()))
	str2 := strconv.Itoa(rand.Intn(100))
	id, _ := strconv.ParseInt(str1+str2, 10, 64)
	return id
}
