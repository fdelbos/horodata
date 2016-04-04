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
		"delbos",
		"about",
		"accept",
		"decline",
		"help",
		"links",
		"my",
		"me",
		"new",
		"nodes",
		"public",
		"references",
		"role",
		"user",
		"users",
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

func Gen(str string, checkFn func(string) (interface{}, error)) (string, error) {
	url := EscapeUrl(str)

	if url == "" {
		return Gen(RandomDigits(4), checkFn)
	}
	if _, err := checkFn(url); err == errors.NotFound {
		return url, nil
	}
	for true {
		test := fmt.Sprintf("%s-%s", url, RandomDigits(6))
		if _, err := checkFn(test); err == errors.NotFound {
			return test, nil
		} else if err != nil {
			return url, err
		}
	}

	return url, nil // should never be called...
}
