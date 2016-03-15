package helpers

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	"fmt"
	"github.com/dchest/uniuri"
	neturl "net/url"
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

func Gen(str string, checkFn func(string) (interface{}, error)) (string, error) {
	url := neturl.QueryEscape(str)
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
