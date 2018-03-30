package main

import (
	_ "github.com/ximply/ngxvts-zabbix/routers"
	"github.com/ximply/ngxvts-zabbix/models"
	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}

