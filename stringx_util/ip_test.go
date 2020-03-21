package stringx_util

import (
	"fmt"
	"github.com/unknwon/com"
	"math/big"
	"testing"
)

func TestHex2Int(t *testing.T) {
	hexString := "0a"

	fmt.Println(com.HexStr2int(hexString))
}

func TestIpv4IntToStr(t *testing.T) {
	ipv4Int := 4758
	fmt.Println(Ipv4IntToStr(ipv4Int))
}

func TestIpv4StrToInt(t *testing.T) {
	ipv4Str := "192.168.1.1"
	fmt.Println(Ipv4StrToInt(ipv4Str))
}

func TestIpv4MaskSize(t *testing.T) {
	fmt.Println(Ipv4MaskSize(24))
}

func TestIpv6IntToStr(t *testing.T) {
	fmt.Println(Ipv6IntToStr(big.NewInt(6553849596)))
}

func TestIpv6StrToInt(t *testing.T) {
	fmt.Println(Ipv6StrToInt("::3223:fe:0"))
}

func TestIpv6MaskSize(t *testing.T) {
	fmt.Println(Ipv6MaskSize(120))
}
