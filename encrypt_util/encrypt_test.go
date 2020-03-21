package encrypt_util

import (
	"fmt"
	"testing"
)

func TestDesEncrypt(t *testing.T) {
	result, _ := DesEncrypt("000000")
	t.Log(result)
}

func TestDesDecrypt(t *testing.T) {
	result, _ := DesDecrypt("7334dd14f48b8e39")
	t.Log(result)
}

func TestAesEncrypt(t *testing.T) {
	result, _ := AesEncrypt("123")
	t.Log(result)
}

func TestAesDecrypt(t *testing.T) {
	result, _ := AesDecrypt("1RiNsqaSuWIac5NhYms/lg==")
	t.Log(result)
}

func TestGenerateSysKey(t *testing.T) {
	sysName := "test"
	sysKey := GenerateSysKey(sysName)

	t.Log(sysName, sysKey)
}

func TestAesEncryptByKeyIv(t *testing.T) {
	clearText := `{"status":1,"pack_type":29,"data":{"game_notice":"\u901a\u77e5\uff1a\u516c\u53f8\u5df2\u5f00\u901a\u5feb\u901f\u5145\u503c\uff08\u5fae\u4fe1\u3001\u652f\u4ed8\u5b9d\uff09\uff0c\u6b22\u8fce\u4f7f\u7528\u3002\u6e29\u99a8\u63d0\u793a\uff1a\u9996\u6b21\u5145\u503c\u8bf7\u6309\u7167\u5145\u503c\u9875\u9762\u63d0\u793a\u64cd\u4f5c\uff01  \u5c0a\u656c\u7684\u5404\u4f4d\u8d35\u5bbe\uff1a\u516c\u53f8\u73b0\u7f51\u5740\u4e3a   www.hjw1688.com  www.hj8828.com  www.hj7878.com   \u901a\u77e5\uff1a\u516c\u53f8\u8fd1\u671f\u53d1\u73b0\u70b9\u51fb\u5ba2\u4eba\u6253\u5bf9\u6ce8\u5957\u6d17\u7801\uff0c\u5982\u53d1\u73b0\u70b9\u51fb\u5ba2\u4eba\u201c\u6076\u610f\u5bf9\u6ce8\u5a31\u4e50\u201d\u516c\u53f8\u5c06\u65e0\u6761\u4ef6\u51bb\u7ed3\u8d26\u53f7\uff0c\u5e76\u6ca1\u6536\u5168\u90e8\u6d17\u7801\u53ca\u4f59\u989d\u3002","hall1_rtmp":"rtmp:\/\/edum.jcx999.com\/live\/thall1","hall1_hls":"http:\/\/eduh.jcx999.com\/live\/thall1\/index.m3u8","hall2_rtmp":"rtmp:\/\/edum.jcx999.com\/live\/thall2","hall2_hls":"rtmp:\/\/edum.jcx999.com\/live\/thall1","balance":"200.00","nickname":"hjtry-038","total_online":14179,"agent_key":0,"gift_chance":0,"system_videos":{"rtmp_address":"rtmp:\/\/edum.jcx999.com\/live","hls_address":"https:\/\/eduh.jcx999.com\/live","flv_address":"https:\/\/flv.jcx999.com\/live","hall1_name":"thall1","hall2_name":"thall2","hd_prefix":"g"}}}`
	rst, err := AesEncryptByKeyIv(clearText, "d8234e682ce9b011", "abcdef0123456789")
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(rst)
}

func TestAesDecryptByKeyIv(t *testing.T) {
	cipherText := `KFv6tZlZbxQJxTLL6kwB1VzAM+Q/7gOdOO22EoGBWEJWBPt4KInThPu3TyLo/tS8IpEmcQiB5/ArblbHKQt9AykeEPWujboNuj/6ulyIOq/Z36RNIG9N1GziYe9eg8D0yhADjj/1ib8KxsJSVjOIdQpTWrc0CimNIQNdGdvYliL0DII5B8gm7lFfcWgjgmCYEIAvIkROFuj0kovIn4HF+MUg6Tqnt6TvKvqs2MPmLB0h7n42E8vThFRaq1RsUqeLvmTYOO4rw3zNigx/ewVPDRyHGBd0GtVgYU+NB1B4MSPjrADKfvc/4ltpxnKh0z/WsXYDLliSUoTft5ZvYO+Xiwdbw14XAVJwMBLSCM1/FiP2fihxzcGp93tCV8s79WoU2z69M17a4egnwryauusyA4lUk0pxpGZauvySNcdgKqGpi0G4+0P7l8SSLTGE7ZIjZIzMc/iBTBrmcuDyFPVPInaflSUOppQANEWkDqlteneCJjvD0hGy8GyWVleRyIrvf+95Y6N4k320olulzJIYNHwY1l5CpYB1thXwdh6WYhyGCElgE2vBaK+NCgy1lQI1GZFHXa++JDCV2oMA6oPCjLqX561n524l1PWsH178pFzVh+PDz9FLbz6mk2jFjcyJGllB3hex3992JZq/rAyJe5ITW/dMNOdriO8I1KQIPfi0Xfj3/5jU3vDx/SpZeaZh6p6T0rv/ReqRXsIfeizc4LME19yHFkd2OD4/m4UgfkEGNsx1U79vbn4N22ZxSAWw/VTyYkXgBn3GTkK669Xgyr1rCajDk3okWxj8HM+I1hVijePaFqNNmXUk41NAXNDsFRdB8karL85ty8Fj6qr0t9914MKTW44ByRb0dLgUUCyNo/Ym20UlUuSO5lSjJl2Y2rE0pXRKlafb+ni8WvHEfyuonnt0x5gKQHkhQmu/tb9OEDfAXsP3X/qKwVCaNDNmU1aGgZiw0G+4zuhJ21e6Mw1SvOzVz0zajqsaQyBf8kM2mYSnvNAcRSoV1+jpJ8sAfzW9l8kIP59hS+G1gXKX0vD6NQACfHht+kbPf+q+bEPL6d/pFE+LQZWKu0ZFZuUnsIodcDwsjlHQ0qu0BNYjiqqBXMwEA7/RnZhSVozRatt/6DTCG2yxlq5ptO5iVY2JQEm4hW4pfVoCwavn3rczG5h2QBfd7eF57eB41rQ4uxPFCHjmwUHVZCADpy3ChrJ3t2VNbcJamY1QEBGGhiYPrJvjoCo4UAXHv6P8azjnBEsa9dGD40qw832ftqrtY47OTszYVb9mzEeZrZ5098FHvBvTfe6lkb2bu3VIZoIbyyMhW17xAW+JweVOxcm2h6CzlNDN/ZJ/3hu/5hkmg0VL0AFZeELgghLzqtK8S8hCsh26+BjjFxY27LoJxZq6bT4Q8OQq5NZu3qlSito7kpCqTwaV/GjCP3ZwDnn5E+dVmiy0XPLYjh6Vp4aX0EJblMUrzDlsoRPl0Ir5RAcz54lUrYwU/TZWpCEeRXxgZLQ6SMxHHAFvcIg/Uj4txdDeLUowJxYPWHezWMW7QPXYj7r+yYAROAwvd5/Ql+X0awjjc9r/++UMhjxagBS8pgHXD7OoawmzOkX7Ov+K8MBnPSQQmqwuK46F4ONO6kZ6FA8OQqQ31MBT/698X5OFIwAt8AR4vyLAfWKQwJ3EJWKCs9jtwHX9ZCb1BZPEWGSNuqDvCn0kptHWzZDKEczMj94Bjd01Ct5sRT7wv7m7gWVvWJFgFJuzTq2SpSK+1V66LQCGooCWxnsU771hnY/LtqmCoB9kfMXJT4+vCBzvPw/bN3Vh4w==`
	rst, err := AesDecryptByKeyIv(cipherText, "d8234e682ce9b011", "abcdef0123456789")
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(rst)
}

func TestMd5(t *testing.T) {
	str := "5588"
	a := Md5(str)
	fmt.Println(a)

	b := Md5(a)
	fmt.Println(b)
}

func Test_1(t *testing.T) {
	a, err := DesDecrypt("db124b3bffdbd3a455f6a8fc48e4202ce1285b1113277d7aa3b95bb7987ac60e0cca66a5e4b017d4f440b7de1488328f898aa9aa75d1132e783f31aaa591bdd713f78b9323eba4b3")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}
