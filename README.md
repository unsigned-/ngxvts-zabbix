ngxvts-zabbix
===========================

Get nginx status(eg. http status code, respone time, traffic) from vozlt/nginx-module-vts to zabbix

## Dependencies
* [nginx](http://nginx.org)
* [nginx-module-vts](https://github.com/vozlt/nginx-module-vts)

## How to use
### setup nginx-module-vts
* [nginx-module-vts-install](https://github.com/vozlt/nginx-module-vts#installation)


### run ngxvts-zabbix
    ./ngxvts-zabbix &


### http api
1. **status update time (nowMsec)**
```Java
   eg. curl http://localhost/status/api/v1/sysinfo/tsmsec or curl http://localhost/status/api/v1/sysinfo/tsmsec?cache=true
```
2. **vhosts(for zabbix discovery)**
```Java
   eg. curl http://localhost/status/api/v1/vhosts
```
3. **connections(connections)**
```Java
   eg. curl http://localhost/status/api/connections/active or curl http://localhost/status/api/connections/active?cache=true
   eg. curl http://localhost/status/api/connections/reading or curl http://localhost/status/api/connections/reading?cache=true
   eg. curl http://localhost/status/api/connections/writing or curl http://localhost/status/api/connections/writing?cache=true
   eg. curl http://localhost/status/api/connections/waiting or curl http://localhost/status/api/connections/waiting?cache=true
   eg. curl http://localhost/status/api/connections/accepted or curl http://localhost/status/api/connections/accepted?cache=true
   eg. curl http://localhost/status/api/connections/handled or curl http://localhost/status/api/connections/handled?cache=true
   eg. curl http://localhost/status/api/connections/requests or curl http://localhost/status/api/connections/requests?cache=true
```
3. **serverzones(serverZones)**

   *:vhost replace with your real vhost name*
```Java
   eg. curl http://localhost/status/api/serverzones/:vhost/requestcounter or curl http://localhost/status/api/serverzones/:vhost/requestcounter?cache=true
   eg. curl http://localhost/status/api/serverzones/:vhost/inbytes or curl http://localhost/status/api/serverzones/:vhost/inbytes?cache=true
   eg. curl http://localhost/status/api/serverzones/:vhost/outbytes or curl http://localhost/status/api/serverzones/:vhost/outbytes?cache=true
   eg. curl http://localhost/status/api/serverzones/:vhost/responses1xx or curl http://localhost/status/api/serverzones/:vhost/responses1xx?cache=true
   eg. curl http://localhost/status/api/serverzones/:vhost/responses2xx or curl http://localhost/status/api/serverzones/:vhost/responses2xx?cache=true
   eg. curl http://localhost/status/api/serverzones/:vhost/responses3xx or curl http://localhost/status/api/serverzones/:vhost/responses3xx?cache=true
   eg. curl http://localhost/status/api/serverzones/:vhost/responses4xx or curl http://localhost/status/api/serverzones/:vhost/responses4xx?cache=true
   eg. curl http://localhost/status/api/serverzones/:vhost/responses5xx or curl http://localhost/status/api/serverzones/:vhost/responses5xx?cache=true
   eg. curl http://localhost/status/api/serverzones/:vhost/requestmsec or curl http://localhost/status/api/serverzones/:vhost/requestmsec?cache=true
   eg. curl http://localhost/status/api/serverzones/:vhost/maxrequestmsec or curl http://localhost/status/api/serverzones/:vhost/maxrequestmsec?cache=true
```
