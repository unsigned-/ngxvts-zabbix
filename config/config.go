package config

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

func NgxvtsKeyAllStatusInfo() string {
	return fmt.Sprintf("%s%s", ngxvtsKeyPrefix(), "all")
}

func NgxvtsKeyAllStatusInfoUpdateTime() string {
	return fmt.Sprintf("%s%s", ngxvtsKeyPrefix(), "updatetime")
}

func ngxvtsKeyPrefix() string {
	return "nginx.status."
}

func NgxvtsCacheEnable() bool {
	return beego.AppConfig.DefaultBool("ngxvts.cacheenable", true)
}

func NgxvtsCacheGcInterval() int {
	return beego.AppConfig.DefaultInt("ngxvts.cachegcinterval", 3600)
}

func NgxVhostsPath() string {
	p := beego.AppConfig.DefaultString("ngxvts.vhostspath ", "/usr/local/nginx/conf/vhosts/")
	if !strings.HasSuffix(p, "/") {
		p += "/"
	}
	return p
}

func NgxVhostsIgnore() []string {
	l := []string{}
	tmp := strings.Split(beego.AppConfig.String("ngxvts.vhostsignore"), "|")
	for _, s := range tmp {
		if len(s) > 0 {
			l = append(l, s)
		}
	}
	return l
}

func NgxZbxVhostVar() string {
	return beego.AppConfig.DefaultString("ngxvts.zbxvhostvar", "NGX_VHOST")
}

func NgxvtsStatusUrl() string {
	return beego.AppConfig.DefaultString("ngxvts.statusurl", "http://localhost:10000/status")
}

func NgxvtsRetry() int {
	return beego.AppConfig.DefaultInt("ngxvts.retry", 3)
}

func NgxvtsUpdateCheckInterval() time.Duration {
	t := beego.AppConfig.DefaultInt("ngxvts.updatecheckinterval", 5)
	return time.Second * time.Duration(t)
}

func NgxvtsUpdateLeftSec() time.Duration {
	t := beego.AppConfig.DefaultInt("ngxvts.updateleftsec", 5)
	return time.Second * time.Duration(t)
}

func NgxvtsTimeout() time.Duration {
	t := beego.AppConfig.DefaultInt("ngxvts.timeout", 3)
	return time.Second * time.Duration(t)
}

func NgxvtsCacheTime() time.Duration {
	t := beego.AppConfig.DefaultInt("ngxvts.cachetimeout", 30)
	return time.Second * time.Duration(t)
}
