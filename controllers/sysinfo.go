package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ximply/ngxvts-zabbix/models"
	"fmt"
)

type SysinfoController struct {
	beego.Controller
}

func (c *SysinfoController) TimeStampMSec() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	c.Ctx.Output.Body([]byte(fmt.Sprintf("%d", models.TimeStampMSec(useCache))))
}
