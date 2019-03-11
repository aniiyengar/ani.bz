
package utils

import (
    "strings"
    "errors"
    "math"
)

const base62chars = (
    "abcdefghijklmnopqrstuvwxyz" +
    "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
    "0123456789" )


func IntToBase62(num uint64) (string, error) {
    result := ""

    if num < 0 {
        return "", errors.New("intToBase62: input cannot be negative")
    }

    if num == 0 {
        return "a", nil
    }

    for num > 0 {
        mod := num % 62
        result = string([]rune(base62chars)[mod]) + result
        num = (num - mod) / 62
    }

    return result, nil
}

func Base62ToInt(str string) (uint64, error) {
    var result uint64 = 0
    strlen := len(str)

    for exp := 0; exp < strlen; exp++ {
        charPos := strlen - 1 - exp
        ix := strings.IndexRune(base62chars, []rune(str)[charPos])
        if ix <= -1 {
            return 0, errors.New("base62ToInt: invalid input string")
        }

        result = result + (uint64(math.Pow(62, float64(exp))) * uint64(ix))
    }

    return result, nil
}

func MagicHashForward(x uint64) uint64 {
    x = (x ^ (x >> 30)) * 13787848793156543929
    x = (x ^ (x >> 27)) * 10723151780598845931
    x = x ^ (x << 31)

    return x
}

func MagicHashBackward(x uint64) uint64 {
    x = (x ^ (x >> 31) ^ (x >> 62)) * 3573116690164977347
    x = (x ^ (x >> 27) ^ (x >> 54)) * 10871156337175269513
    x = x ^ (x >> 30) ^ (x >> 60);

    return x
}
