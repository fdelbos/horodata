package helpers

import (
	"fmt"
	"strings"

	"dev.hyperboloide.com/fred/horodata/models/errors"
	"github.com/dchest/uniuri"
)

var (
	hexChars  = []byte("abcdef0123456789")
	digitChar = []byte("0123456789")
	Reserved  = []string{
		"abonnement",
		"about",
		"accept",
		"billing",
		"decline",
		"delbos",
		"frederic",
		"fred",
		"help",
		"horodata",
		"links",
		"my",
		"me",
		"new",
		"nodes",
		"public",
		"references",
		"role",
		"static",
		"user",
		"users",
		"usinedata",
		"visit",
	}
)

func RandomHex(n int) string {
	return uniuri.NewLenChars(n, hexChars)
}

func RandomDigits(n int) string {
	return uniuri.NewLenChars(n, digitChar)
}

var replacer = strings.NewReplacer(
	"<", "-", ">", "-", "#", "-", "%", "-", "\"", "-", "{", "-", "}", "-", "|", "-", "\\", "-", "^", "-", "[", "-", "]", "-", "`", "-", "?", "-", "&", "-", "=", "-", "+", "-", "/", "-", " ", "-")

func EscapeUrl(str string) string {
	return replacer.Replace(str)
}

func IsReserved(str string) bool {
	for _, r := range Reserved {
		if str == r {
			return true
		}
	}
	return false
}

func Gen(str string, checkFn func(string) (interface{}, error)) (string, error) {
	url := EscapeUrl(str)

	if url == "" {
		return Gen(RandomDigits(4), checkFn)
	} else if IsReserved(url) {
		return Gen(fmt.Sprintf("%s-%s", url, RandomDigits(4)), checkFn)
	} else if _, err := checkFn(url); err == errors.NotFound {
		return url, nil
	}
	for true {
		test := fmt.Sprintf("%s-%s", url, RandomDigits(4))
		if _, err := checkFn(test); err == errors.NotFound {
			return test, nil
		} else if err != nil {
			return url, err
		}
	}

	return url, nil // should never be called...
}
