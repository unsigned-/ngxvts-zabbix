package models

import (
	"github.com/parnurzeal/gorequest"
	"github.com/json-iterator/go"
	"github.com/ximply/ngxvts-zabbix/config"
	"time"
	"net/http"
	"fmt"
)

type NginxVts struct {
	HostName     string `json:"hostName"`
	NginxVersion string `json:"nginxVersion"`
	LoadMsec     int64  `json:"loadMsec"`
	NowMsec      int64  `json:"nowMsec"`
	Connections  Conns  `json:"connections"`
	ServerZones   map[string]Server              `json:"serverZones"`
	UpstreamZones map[string][]Upstream          `json:"upstreamZones"`
	FilterZones   map[string]map[string]Upstream `json:"filterZones"`
	CacheZones    map[string]Cache               `json:"cacheZones"`
}

type Conns struct {
	Active   uint64 `json:"active"`
	Reading  uint64 `json:"reading"`
	Writing  uint64 `json:"writing"`
	Waiting  uint64 `json:"waiting"`
	Accepted uint64 `json:"accepted"`
	Handled  uint64 `json:"handled"`
	Requests uint64 `json:"requests"`
}

type Server struct {
	RequestCounter uint64 `json:"requestCounter"`
	InBytes        uint64 `json:"inBytes"`
	OutBytes       uint64 `json:"outBytes"`
	RequestMsec    float64 `json:"requestMsec"`
	RequestMsecs   struct{
		Times []uint64 `json:"times"`
		Msecs []float64 `json:"msecs"`
	} `json:"requestMsecs"`
	Responses      struct {
		OneXx       uint64 `json:"1xx"`
		TwoXx       uint64 `json:"2xx"`
		ThreeXx     uint64 `json:"3xx"`
		FourXx      uint64 `json:"4xx"`
		FiveXx      uint64 `json:"5xx"`
		Miss        uint64 `json:"miss"`
		Bypass      uint64 `json:"bypass"`
		Expired     uint64 `json:"expired"`
		Stale       uint64 `json:"stale"`
		Updating    uint64 `json:"updating"`
		Revalidated uint64 `json:"revalidated"`
		Hit         uint64 `json:"hit"`
		Scarce      uint64 `json:"scarce"`
	} `json:"responses"`
	OverCounts struct {
		MaxIntegerSize float64 `json:"maxIntegerSize"`
		RequestCounter uint64  `json:"requestCounter"`
		InBytes        uint64  `json:"inBytes"`
		OutBytes       uint64  `json:"outBytes"`
		OneXx          uint64  `json:"1xx"`
		TwoXx          uint64  `json:"2xx"`
		ThreeXx        uint64  `json:"3xx"`
		FourXx         uint64  `json:"4xx"`
		FiveXx         uint64  `json:"5xx"`
		Miss           uint64  `json:"miss"`
		Bypass         uint64  `json:"bypass"`
		Expired        uint64  `json:"expired"`
		Stale          uint64  `json:"stale"`
		Updating       uint64  `json:"updating"`
		Revalidated    uint64  `json:"revalidated"`
		Hit            uint64  `json:"hit"`
		Scarce         uint64  `json:"scarce"`
	} `json:"overCounts"`
}

type Upstream struct {
	Server         string `json:"server"`
	RequestCounter uint64 `json:"requestCounter"`
	InBytes        uint64 `json:"inBytes"`
	OutBytes       uint64 `json:"outBytes"`
	Responses      struct {
		OneXx   uint64 `json:"1xx"`
		TwoXx   uint64 `json:"2xx"`
		ThreeXx uint64 `json:"3xx"`
		FourXx  uint64 `json:"4xx"`
		FiveXx  uint64 `json:"5xx"`
	} `json:"responses"`
	ResponseMsec float64 `json:"responseMsec"`
	RequestMsec  float64 `json:"requestMsec"`
	Weight       uint64 `json:"weight"`
	MaxFails     uint64 `json:"maxFails"`
	FailTimeout  uint64 `json:"failTimeout"`
	Backup       bool   `json:"backup"`
	Down         bool   `json:"down"`
	OverCounts   struct {
		MaxIntegerSize float64 `json:"maxIntegerSize"`
		RequestCounter uint64  `json:"requestCounter"`
		InBytes        uint64  `json:"inBytes"`
		OutBytes       uint64  `json:"outBytes"`
		OneXx          uint64  `json:"1xx"`
		TwoXx          uint64  `json:"2xx"`
		ThreeXx        uint64  `json:"3xx"`
		FourXx         uint64  `json:"4xx"`
		FiveXx         uint64  `json:"5xx"`
	} `json:"overCounts"`
}

type Cache struct {
	MaxSize   uint64 `json:"maxSize"`
	UsedSize  uint64 `json:"usedSize"`
	InBytes   uint64 `json:"inBytes"`
	OutBytes  uint64 `json:"outBytes"`
	Responses struct {
		Miss        uint64 `json:"miss"`
		Bypass      uint64 `json:"bypass"`
		Expired     uint64 `json:"expired"`
		Stale       uint64 `json:"stale"`
		Updating    uint64 `json:"updating"`
		Revalidated uint64 `json:"revalidated"`
		Hit         uint64 `json:"hit"`
		Scarce      uint64 `json:"scarce"`
	} `json:"responses"`
	OverCounts struct {
		MaxIntegerSize float64 `json:"maxIntegerSize"`
		InBytes        uint64  `json:"inBytes"`
		OutBytes       uint64  `json:"outBytes"`
		Miss           uint64  `json:"miss"`
		Bypass         uint64  `json:"bypass"`
		Expired        uint64  `json:"expired"`
		Stale          uint64  `json:"stale"`
		Updating       uint64  `json:"updating"`
		Revalidated    uint64  `json:"revalidated"`
		Hit            uint64  `json:"hit"`
		Scarce         uint64  `json:"scarce"`
	} `json:"overCounts"`
}

var g_checking = false

func nginxSatusInfo(retry int, timeout time.Duration) string {
	statusUrl := config.NgxvtsStatusUrl()
	req := gorequest.New()
	_, body, errs := req.Retry(retry, timeout,
		http.StatusBadRequest, http.StatusInternalServerError).Get(statusUrl).End()
	if errs != nil {
		return ""
	}
	return body
}

func updateStatusInfoToCache() interface{} {
	content := nginxSatusInfo(config.NgxvtsRetry(), config.NgxvtsTimeout())
	ngvts := NginxVts{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal([]byte(content), &ngvts)
	if err != nil {
		CacheSet(config.NgxvtsKeyAllStatusInfoUpdateTime(), time.Now(), config.NgxvtsCacheTime())
		return nil
	}
	CacheSet(config.NgxvtsKeyAllStatusInfoUpdateTime(), time.Now(), 0)
	CacheSet(config.NgxvtsKeyAllStatusInfo(), ngvts, config.NgxvtsCacheTime())
	return ngvts
}

func nginxStatusInfoUpdateCheck() {
	cacheTime := config.NgxvtsCacheTime()
	checkInterval := config.NgxvtsUpdateCheckInterval()
	ticker := time.NewTicker(checkInterval)
	left := config.NgxvtsUpdateLeftSec()
	for _ = range ticker.C {
		if g_checking == false {
			g_checking = true
			updateTime := CacheGet(config.NgxvtsKeyAllStatusInfoUpdateTime(), time.Now())
			futureTime := updateTime.Content.(time.Time).Add(cacheTime)
			now := time.Now()
			if futureTime.Sub(now) < left {
				updateStatusInfoToCache()
			}
			g_checking = false
		}
	}
}