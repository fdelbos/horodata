package cache

import (
	"bitbucket.com/hyperboloide/horo/models/errors"
	"encoding/json"
	"fmt"
	"gopkg.in/redis.v3"
	"time"
)

func Ping() error {
	return client.Ping().Err()
}

func mkPakage(pkg, key string) string {
	return fmt.Sprintf("%s/%s", pkg, key)
}

func Set(key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(key, string(data[:]), expiration).Err()
}

func SetPackage(pkg, key string, value interface{}, expiration time.Duration) error {
	return Set(mkPakage(pkg, key), value, expiration)
}

func Get(key string, obj interface{}) error {
	v, err := client.Get(key).Result()
	if err == redis.Nil {
		return errors.NotFound
	} else if err != nil {
		return err
	}
	return json.Unmarshal([]byte(v), obj)
}

func GetPackage(pkg, key string, obj interface{}) error {
	return Get(mkPakage(pkg, key), obj)
}

func Del(key string) error {
	_, err := client.Del(key).Result()
	return err
}

func DelPackage(pkg, key string) error {
	return Del(mkPakage(pkg, key))
}
