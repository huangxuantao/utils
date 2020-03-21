package net_util

import (
	"net"
	"strings"
)

func GetAllAddress() string {
	var address string
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return address
	}
	var addressList []string
	for _, addr := range addrList {
		addressList = append(addressList, addr.String())
	}
	address = strings.Join(addressList, ",")
	return address
}
