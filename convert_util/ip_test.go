package convert_util

import (
	"fmt"
	"testing"
)

func TestIpv4StringToInt(t *testing.T) {
	ipv4String := "255.255.255.255"
	fmt.Println(Ipv4StringToInt(ipv4String))
}

func TestIpv4IntToString(t *testing.T) {
	ipv4Int := 374953
	fmt.Println(Ipv4IntToString(ipv4Int))
}
