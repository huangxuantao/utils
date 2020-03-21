package convert_util

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"math/big"
	"net"
)

var (
	bigInt0    = big.NewInt(0)
	ipv4IntMax = 4294967295 // Ipv4 int 的最大值 255.255.255.255
)

// ipv4 将 xxx.xxx.xxx.xxx 转换成 int
func Ipv4StringToInt(ipv4String string) (int, bool) {
	if !govalidator.IsIPv4(ipv4String) {
		// 非ipv4格式，直接返回错误
		return int(bigInt0.Int64()), false
	}
	ipv4Int := big.NewInt(0)
	ipv4Int.SetBytes(net.ParseIP(ipv4String).To4())
	return int(ipv4Int.Int64()), true
}

// ipv4 将 int 转换为 xxx.xxx.xxx.xxx
func Ipv4IntToString(ipv4Int int) (string, bool) {
	if ipv4Int < 0 || ipv4Int > ipv4IntMax {
		return "", false
	}
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ipv4Int>>24), byte(ipv4Int>>16), byte(ipv4Int>>8), byte(ipv4Int)), true
}
