package models

import (
	"github.com/astaxie/beego/cache"
	"github.com/ximply/ngxvts-zabbix/config"
	"fmt"
)

var ngxvtsCache interface{}
var ngxvtsErr error

func Init() {
	if config.NgxvtsCacheEnable() {
		ngxvtsCache, ngxvtsErr = cache.NewCache("memory",
			fmt.Sprintf("`{\"interval\":%d}`", config.NgxvtsCacheGcInterval()))
		if ngxvtsErr != nil {
			panic(ngxvtsErr.Error())
		}
		CacheSet("ngxvts.cacheenable", true, 0)

		updateStatusInfoToCache()
		go nginxStatusInfoUpdateCheck()
	}
}


