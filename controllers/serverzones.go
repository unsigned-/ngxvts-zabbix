package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ximply/ngxvts-zabbix/models"
	"fmt"
)

type ServerZonesController struct {
	beego.Controller
}

func (c *ServerZonesController) RequestCounter(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.RequestCounter(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) InBytes(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.InBytes(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) OutBytes(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.OutBytes(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses1xx(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Responses1xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses2xx(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Responses2xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses3xx(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Responses3xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses4xx(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Responses4xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) Responses5xx(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.Responses5xx(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%d", val.(uint64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) RequestMsec(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.RequestMsec(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%.1f", val.(float64))
	}
	c.Ctx.Output.Body([]byte(resp))
}

func (c *ServerZonesController) MaxRequestMsec(vhost string) {
	useCache := true
	c.Ctx.Input.Bind(&useCache, "cache")
	val := models.MaxRequestMsec(vhost, useCache)
	resp := ""
	if val != nil {
		resp = fmt.Sprintf("%.1f", val.(float64))
	}
	c.Ctx.Output.Body([]byte(resp))
}