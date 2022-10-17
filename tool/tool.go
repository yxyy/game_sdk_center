package tool

import (
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
