package time_util

import (
	"fmt"
	"testing"
)

func TestGetNowTime(t *testing.T) {
	t.Log(GetNowTime())
}

func TestGetTimeByFormatString(t *testing.T) {
	format := "02/Jan/2006:15:04:05 +0800"
	timeString := "07/Feb/2019:14:40:59 +0800"
	timeLocation, err := GetTimeByFormatString(format, timeString)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(timeLocation.Unix(), timeLocation.String())
}
