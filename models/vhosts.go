package models

import (
	"github.com/ximply/ngxvts-zabbix/utils"
)

type VhostItem struct {
	VhostList []string
	VhostZabbixDiscoveryStr string
}

func VhostList() VhostItem {
	s, l := utils.DiscoeryNginxServerName()
	vi := VhostItem{
		VhostList: l,
		VhostZabbixDiscoveryStr: s,
	}
	return vi
}


