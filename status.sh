#/bin/sh

curl="/usr/bin/curl -s"
apiUrlPrefix="http://localhost:10002/api/v1"
model=$1
item=$2
cache=$3
vhost=$4

case $model in
        sysinfo)
                case $item in
                        tsmsec)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/sysinfo/tsmsec?cache=true
                                else
                                        $curl $apiUrlPrefix/sysinfo/tsmsec
                                fi
                                ;;
                esac
                ;;

        vhosts)
        $curl $apiUrlPrefix/vhosts
        ;;

        connections)
                case $item in
                        active)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/connections/active?cache=true
                                else
                                        $curl $apiUrlPrefix/connections/active
                                fi
                                ;;

                        reading)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/connections/reading?cache=true
                                else
                                        $curl $apiUrlPrefix/connections/reading
                                fi
                                ;;

                        writing)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/connections/writing?cache=true
                                else
                                        $curl $apiUrlPrefix/connections/writing
                                fi
                                ;;

                        waiting)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/connections/waiting?cache=true
                                else
                                        $curl $apiUrlPrefix/connections/waiting
                                fi
                                ;;

                        accepted)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/connections/accepted?cache=true
                                else
                                        $curl $apiUrlPrefix/connections/accepted
                                fi
                                ;;

                        handled)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/connections/handled?cache=true
                                else
                                        $curl $apiUrlPrefix/connections/handled
                                fi
                                ;;

                        requests)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/connections/requests?cache=true
                                else
                                        $curl $apiUrlPrefix/connections/requests
                                fi
                                ;;                                                                
                esac
                ;;

        serverzones)
                case $item in
                        requestcounter)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/requestcounter?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/requestcounter
                                fi
                                ;;

                        inbytes)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/inbytes?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/inbytes
                                fi
                                ;; 

                        inbytes)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/inbytes?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/inbytes
                                fi
                                ;; 

                        inbytes)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/inbytes?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/inbytes
                                fi
                                ;; 

                        outbytes)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/outbytes?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/outbytes
                                fi
                                ;; 

                        responses1xx)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses1xx?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses1xx
                                fi
                                ;;

                        responses2xx)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses2xx?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses2xx
                                fi
                                ;;  

                        responses3xx)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses3xx?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses3xx
                                fi
                                ;;  

                        responses4xx)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses4xx?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses4xx
                                fi
                                ;;  

                        responses5xx)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses5xx?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/responses5xx
                                fi
                                ;; 

                        requestmsec)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/requestmsec?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/requestmsec
                                fi
                                ;; 

                        maxrequestmsec)
                                if [ $cache -eq 1 ]
                                then
                                        $curl $apiUrlPrefix/serverzones/$vhost/maxrequestmsec?cache=true
                                else
                                        $curl $apiUrlPrefix/serverzones/$vhost/maxrequestmsec
                                fi
                                ;;                                                                                                                                                                                                                                                                                                                                                                  
                esac
                ;;
esac
