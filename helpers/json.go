package helpers

import (
	"encoding/json"
)

func InterfaceToJsonStr(i interface{}) (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}

func JsonStrToInterface(str string, i interface{}) error {
	b := []byte(str)
	return json.Unmarshal(b, i)
}
