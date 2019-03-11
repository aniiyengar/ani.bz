
package utils

import (
    "strings"
    "strconv"
    "errors"
    "math"
    "os"
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
    modEnv := os.Getenv("ANI_BZ_MOD_VARS")
    envStrings := strings.Split(modEnv, " ")
    modPInt, _ := strconv.ParseInt(envStrings[0], 10, 64)
    modNInt, _ := strconv.ParseInt(envStrings[2], 10, 64)

    modP := uint64(modPInt)
    modN := uint64(modNInt)

    return (x * modP) % modN
}

func MagicHashBackward(x uint64) uint64 {
    modEnv := os.Getenv("ANI_BZ_MOD_VARS")
    envStrings := strings.Split(modEnv, " ")
    modQInt, _ := strconv.ParseInt(envStrings[1], 10, 64)
    modNInt, _ := strconv.ParseInt(envStrings[2], 10, 64)

    modQ := uint64(modQInt)
    modN := uint64(modNInt)

    return (x * modQ) % modN
}
