package stringx_util

import (
	"fmt"
	"testing"
)

func TestListStringAddQuotation(t *testing.T) {
	s := []string{
		"aaa", "bbb", "ccc",
	}
	fmt.Println(ListStringAddQuotation(s))
}

func TestSeparateByComma(t *testing.T) {
	s := "aaa,bbb,ccc"
	fmt.Println(SeparateByComma(s))
}
