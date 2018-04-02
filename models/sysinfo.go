package models

import (
	"github.com/json-iterator/go"
	"github.com/ximply/ngxvts-zabbix/config"
)

func TimeStampMSec(useCache bool) interface{} {
	if !config.NgxvtsCacheEnable() || !useCache {
		content := nginxSatusInfo(config.NgxvtsRetry(), config.NgxvtsTimeout())
		if len(content) == 0 {
			return nil
		}

		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		ngvts := NginxVts{}
		err := json.Unmarshal([]byte(content), &ngvts)
		if err != nil {
			return nil
		}

		return ngvts.NowMsec
	}

	v := CacheGet(config.NgxvtsKeyAllStatusInfo(), nil)
	if v.Content == nil {
		return nil
	}
	return v.Content.(NginxVts).NowMsec
}

