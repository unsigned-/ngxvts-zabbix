package models

import (
	"github.com/ximply/ngxvts-zabbix/utils"
	"github.com/ximply/ngxvts-zabbix/config"
	"github.com/json-iterator/go"
)

func serverZones(vhost string, useCache bool) interface{} {
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
		return ngvts.ServerZones[vhost]
	}

	v := CacheGet(config.NgxvtsKeyAllStatusInfo(), nil)
	if v.Content == nil {
		return nil
	}
	return v.Content.(NginxVts).ServerZones[vhost]
}

func RequestCounter(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}
	return content.(Server).RequestCounter
}

func InBytes(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}
	return content.(Server).InBytes
}

func OutBytes(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}
	return content.(Server).OutBytes
}

func Responses1xx(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}
	return content.(Server).Responses.OneXx
}

func Responses2xx(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}
	return content.(Server).Responses.TwoXx
}

func Responses3xx(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}
	return content.(Server).Responses.ThreeXx
}

func Responses4xx(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}
	return content.(Server).Responses.FourXx
}

func Responses5xx(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}
	return content.(Server).Responses.FiveXx
}

func RequestMsec(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}
	return content.(Server).RequestMsec
}

func MaxRequestMsec(vhost string, useCache bool) interface{} {
	content := serverZones(vhost, useCache)
	if content == nil {
		return nil
	}

	arr := content.(Server).RequestMsecs.Msecs
	if len(arr) == 0 {
		return nil
	}

	utils.Float64s(arr)
	return arr[len(arr) - 1]
}