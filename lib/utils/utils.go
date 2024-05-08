// generic functions

package go_utils

import (
	"math/rand"
	"time"
)

// random pick from array
func RandFromArray[T any](array []T) T {
    return array[rand.Intn(len(array))]
}

// random between range, inclusive of max
func RandRange(min, max int) int {
    return rand.Intn(max+1-min)+min
}

// get a basic timestamp
func GetTimeStamp() string {
    return time.Now().Format("15:04:05")
}