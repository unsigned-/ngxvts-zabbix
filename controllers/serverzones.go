package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ximply/ngxvts-zabbix/models"
	"fmt"
)

type ServerZonesController struct {
	beego.Controller
}

func (c *ServerZonesController) RequestCounter() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.RequestCounter(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) InBytes() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.InBytes(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) OutBytes() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.OutBytes(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses1xx() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.Responses1xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses2xx() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.Responses2xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses3xx() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.Responses3xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses4xx() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.Responses4xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses5xx() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.Responses5xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) RequestMsec() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.RequestMsec(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%.1f", val.(float64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) MaxRequestMsec() {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	vhost := c.Ctx.Input.Param(":vhost")
	val := models.MaxRequestMsec(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%.1f", val.(float64))
	}
	c.Ctx.Output.Body([]byte(resp))
}