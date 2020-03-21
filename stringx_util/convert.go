package stringx_util

import (
	"fmt"
	"strings"
)

// 给字符串列表的每个元素加上引号
func ListStringAddQuotation(src []string) []string {
	var rst []string
	for _, item := range src {
		str := fmt.Sprintf("'%s'", item)
		rst = append(rst, str)
	}
	return rst
}

// 通过逗号分隔
func SeparateByComma(src string) []string {
	var rst []string
	raw := strings.Split(src, ",")
	for _, item := range raw {
		if item != "" {
			rst = append(rst, item)
		}
	}
	return rst
}
