package models

import (
	"github.com/astaxie/beego/cache"
	"time"
)

type CacheItem struct {
	Key string
	Content interface{}
	Exist bool
}

func CacheGet(key string, defaultVal interface{}) CacheItem {
	ci := CacheItem{
		Key: key,
		Content: defaultVal,
		Exist: false,
	}
	if len(key) == 0 {
		return ci
	}

	v := ngxvtsCacheGet(key)
	if v == nil {
		return ci
	}

	ci.Content = v
	ci.Exist = true
	return ci
}

func CacheGetString(key string, defaultVal string) CacheItem {
	ci := CacheItem{
		Key: key,
		Content: defaultVal,
		Exist: false,
	}
	if len(key) == 0 {
		return ci
	}

	v := ngxvtsCacheGet(key)
	if v == nil {
		return ci
	}

	ci.Content = v.(string)
	ci.Exist = true
	return ci
}

func CacheGetInt64(key string, defaultVal int64) CacheItem {
	ci := CacheItem{
		Key: key,
		Content: defaultVal,
		Exist: false,
	}
	if len(key) == 0 {
		return ci
	}

	v := ngxvtsCacheGet(key)
	if v == nil {
		return ci
	}

	ci.Content = v.(int64)
	ci.Exist = true
	return ci
}

func CacheGetBool(key string, defaultVal bool) CacheItem {
	ci := CacheItem{
		Key: key,
		Content: defaultVal,
		Exist: false,
	}
	if len(key) == 0 {
		return ci
	}

	v := ngxvtsCacheGet(key)
	if v == nil {
		return ci
	}

	ci.Content = v.(bool)
	ci.Exist = true
	return ci
}

func ngxvtsCacheGet(key string) interface{} {
	if ngxvtsCache != nil {
		return ngxvtsCache.(cache.Cache).Get(key)
	}
	return nil
}

func CacheSet(key string, val interface{}, timeout time.Duration) (error, bool) {
	if len(key) == 0 {
		return nil, false
	}
	return ngxvtsCacheSet(key, val, timeout)
}

func ngxvtsCacheSet(key string, val interface{}, timeout time.Duration) (error, bool) {
	if ngxvtsCache != nil {
		return ngxvtsCache.(cache.Cache).Put(key, val, timeout), true

	}
	return nil, false
}
