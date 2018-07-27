
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


func IntToBase62(num int) (string, error) {
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

func Base62ToInt(str string) (int, error) {
    result := 0
    strlen := len(str)

    for exp := 0; exp < strlen; exp++ {
        charPos := strlen - 1 - exp
        ix := strings.IndexRune(base62chars, []rune(str)[charPos])
        if ix <= -1 {
            return -1, errors.New("base62ToInt: invalid input string")
        }

        result = result + (int(math.Pow(62, float64(exp))) * ix)
    }

    return result, nil
}
