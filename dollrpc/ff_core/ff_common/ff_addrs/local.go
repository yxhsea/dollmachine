package ff_addrs

import (
	"fmt"
	"net"
)

//获取本地局域网IP
func GetLocalAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err.Error())
	}
	var Ip string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				Ip = ipnet.IP.String()
			}
		}
	}
	return Ip
}
