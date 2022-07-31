package util

import (
	"math/rand"
	"time"
)

func MakeRandomStr(digit int) string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, digit)

    if _, err := rand.Read(b); err != nil {
        return ""
    }
    
    var result string
    for _, v := range b {
        // index が letters の長さに収まるように調整
        result += string(letters[int(v)%len(letters)])
    }
    return result
}

func MakeRandomInt() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Int()
}

func MakeRandomBool() bool {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(2) == 0
}