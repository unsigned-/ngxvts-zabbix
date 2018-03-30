package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/ximply/ngxvts-zabbix/models"
)

type ConnectionsController struct {
	beego.Controller
}

func (c *ConnectionsController) Active() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Active(useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ConnectionsController) Reading() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Reading(useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ConnectionsController) Writing() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Writing(useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ConnectionsController) Waiting() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Waiting(useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ConnectionsController) Accepted() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Accepted(useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ConnectionsController) Handled() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Handled(useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ConnectionsController) Requests() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Requests(useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

