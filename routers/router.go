package routers

import (
	"github.com/ximply/ngxvts-zabbix/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //sysinfo
	beego.Router("/api/v1/sysinfo/tsmsec",
		&controllers.SysinfoController{}, "get:TimeStampMSec")

	// vhosts
    beego.Router("/api/v1/vhosts",
    	&controllers.VhostsController{}, "get:Vhosts")

	// connections
	beego.Router("/api/v1/connections/active",
		&controllers.ConnectionsController{}, "get:Active")
	beego.Router("/api/v1/connections/reading",
		&controllers.ConnectionsController{}, "get:Reading")
	beego.Router("/api/v1/connections/writing",
		&controllers.ConnectionsController{}, "get:Writing")
	beego.Router("/api/v1/connections/waiting",
		&controllers.ConnectionsController{}, "get:Waiting")
	beego.Router("/api/v1/connections/accepted",
		&controllers.ConnectionsController{}, "get:Accepted")
	beego.Router("/api/v1/connections/handled",
		&controllers.ConnectionsController{}, "get:Handled")
	beego.Router("/api/v1/connections/requests",
		&controllers.ConnectionsController{}, "get:Requests")

	// serverzones
	beego.Router("/api/v1/serverzones/:vhost/requestcounter",
		&controllers.ServerZonesController{}, "get:RequestCounter")
	beego.Router("/api/v1/serverzones/:vhost/inbytes",
		&controllers.ServerZonesController{}, "get:InBytes")
	beego.Router("/api/v1/serverzones/:vhost/outbytes",
		&controllers.ServerZonesController{}, "get:OutBytes")
	beego.Router("/api/v1/serverzones/:vhost/responses1xx",
		&controllers.ServerZonesController{}, "get:Responses1xx")
	beego.Router("/api/v1/serverzones/:vhost/responses2xx",
		&controllers.ServerZonesController{}, "get:Responses2xx")
	beego.Router("/api/v1/serverzones/:vhost/responses3xx",
		&controllers.ServerZonesController{}, "get:Responses3xx")
	beego.Router("/api/v1/serverzones/:vhost/responses4xx",
		&controllers.ServerZonesController{}, "get:Responses4xx")
	beego.Router("/api/v1/serverzones/:vhost/responses5xx",
		&controllers.ServerZonesController{}, "get:Responses5xx")
	beego.Router("/api/v1/serverzones/:vhost/requestmsec",
		&controllers.ServerZonesController{}, "get:RequestMsec")
	beego.Router("/api/v1/serverzones/:vhost/maxrequestmsec",
		&controllers.ServerZonesController{}, "get:MaxRequestMsec")
}

