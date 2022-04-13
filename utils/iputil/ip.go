package iputil

import (
	"errors"
	"net"
	"strings"
)

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && strings.HasPrefix(ipnet.IP.String(), "192.168") {
			return ipnet.IP.String(), nil
		}
	}

	return "", errors.New("no valid ip")
}
