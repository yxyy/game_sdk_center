package tool

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Uuid() string {

	return time.Now().Format("20060102150405") + strconv.Itoa(Range())

}

func Range() int {

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(8999) + 1000
}

func Salt() string {

	return fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(int(time.Now().Unix()))+strconv.Itoa(Range()))))

}
