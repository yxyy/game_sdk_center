package tool

import (
	"math/rand"
	"time"
)

func Uuid() string {

	return time.Now().Format("20060102150405")

}

func Range() {

	rand.Seed(time.Now().UnixNano())

	rand.Int()
}
