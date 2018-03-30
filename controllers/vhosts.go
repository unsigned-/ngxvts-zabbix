package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ximply/ngxvts-zabbix/models"
)

type VhostsController struct {
	beego.Controller
}

func (c *VhostsController) Vhosts() {
	c.Ctx.Output.Body([]byte(models.VhostList().VhostZabbixDiscoveryStr))
}
