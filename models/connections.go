package models

import (
	"github.com/json-iterator/go"
	"github.com/ximply/ngxvts-zabbix/config"
)

func connections(useCache bool) interface{} {
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
		return ngvts.Connections
	}

	v := CacheGet(config.NgxvtsKeyAllStatusInfo(), nil)
	if v.Content == nil {
		return nil
	}
	return v.Content.(NginxVts).Connections
}

func Active(useCache bool) interface{} {
	content := connections(useCache)
	if content == nil {
		return nil
	}
	return content.(Conns).Active
}

func Reading(useCache bool) interface{} {
	content := connections(useCache)
	if content == nil {
		return nil
	}
	return content.(Conns).Reading
}

func Writing(useCache bool) interface{} {
	content := connections(useCache)
	if content == nil {
		return nil
	}
	return content.(Conns).Writing
}

func Waiting(useCache bool) interface{} {
	content := connections(useCache)
	if content == nil {
		return nil
	}
	return content.(Conns).Waiting
}

func Accepted(useCache bool) interface{} {
	content := connections(useCache)
	if content == nil {
		return nil
	}
	return content.(Conns).Accepted
}

func Handled(useCache bool) interface{} {
	content := connections(useCache)
	if content == nil {
		return nil
	}
	return content.(Conns).Handled
}

func Requests(useCache bool) interface{} {
	content := connections(useCache)
	if content == nil {
		return nil
	}
	return content.(Conns).Requests
}
