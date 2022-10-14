package tool

import (
	"math/rand"
	"time"
)

func Uuid() string {

	return time.Now().Format("20060102150405")

}

func Range() int {

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(9999) + 1000
}
