package stringx_util

import (
	"fmt"
	"github.com/unknwon/com"
	"math"
	"math/big"
	"net"
	"strconv"
	"strings"
)

const (
	Ipv4MaskMax int = 32
	Ipv6MaskMax     = 128
)

// IPv4 xxx.xxx.xxx.xxx 字符串转换为 int
func Ipv4StrToInt(ipv4Str string) (ipv4Int int) {
	bits := strings.Split(ipv4Str, ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	ipv4Int += int(b0) << 24
	ipv4Int += int(b1) << 16
	ipv4Int += int(b2) << 8
	ipv4Int += int(b3)

	return ipv4Int
}

func Ipv4IntToStr(ipv4Int int) (ipv4Str string) {
	ipv4Str = fmt.Sprintf("%d.%d.%d.%d",
		byte(ipv4Int>>24), byte(ipv4Int>>16), byte(ipv4Int>>8), byte(ipv4Int))
	return ipv4Str
}

func Ipv4MaskSize(mask int) int {
	return int(math.Pow(2, float64(Ipv4MaskMax-mask)))
}

func Ipv6StrToInt(ipv6Str string) (ipv6BigInt *big.Int) {
	var ipv6List []string
	bits := strings.Split(ipv6Str, ":")
	if bits[0] == "" {
		bits[0] = "0"
	}
	if bits[len(bits)-1] == "" {
		bits[len(bits)-1] = "0"
	}

	for i := 0; i < len(bits); i++ {
		if bits[i] != "" {
			ipv6List = append(ipv6List, bits[i])
		} else {
			for ii := 0; ii < 8-len(bits)+1; ii++ {
				ipv6List = append(ipv6List, "0")
			}
		}
	}

	ipv6BigInt = big.NewInt(0)
	for i, bit := range ipv6List {
		value, _ := com.HexStr2int(bit)
		valueBI := big.NewInt(int64(value))
		for ii := 0; ii < 8-1-i; ii++ {
			valueBI.Mul(valueBI, big.NewInt(65536))
		}
		ipv6BigInt = ipv6BigInt.Add(ipv6BigInt, valueBI)
	}
	return ipv6BigInt
}

func Ipv6IntToStr(ipv6BigInt *big.Int) (ipv6Str string) {
	var bit [8]*big.Int
	for i := 0; i < 8; i++ {
		bit[i] = big.NewInt(0)
		bit[i].Set(ipv6BigInt)
		for ii := 0; ii < 8-1-i; ii++ {
			bit[i].Div(bit[i], big.NewInt(65536))
		}
		bit[i].Mod(bit[i], big.NewInt(65536))
	}

	ipv6Str = fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x", bit[0], bit[1], bit[2], bit[3], bit[4], bit[5], bit[6], bit[7])
	return net.ParseIP(ipv6Str).String()
}

func Ipv6MaskSize(mask int) *big.Int {
	sum := big.NewInt(1)
	for i := 0; i < Ipv6MaskMax-mask; i++ {
		sum = sum.Mul(sum, big.NewInt(2))
	}
	return sum
}
